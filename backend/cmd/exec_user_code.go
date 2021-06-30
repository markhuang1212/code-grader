package cmd

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/markhuang1212/code-grader/types"
	"github.com/pkg/errors"
)

const imageExec = "markhuang1212/code-grader/runtime-exec:latest"

func exec_user_code(ctx context.Context, gr types.GradeRequest) ([]byte, error) {

	var testCase types.TestCaseOptions

	testCaseJson, err := os.ReadFile(filepath.Join(AppRoot, "testcases", gr.TestCaseName, "testcase.json"))
	if err != nil {
		return nil, errors.Wrap(InternalError, "cannot open testcase.json")
	}

	err = json.Unmarshal(testCaseJson, &testCase)
	if err != nil {
		return nil, errors.Wrap(InternalError, "cannot parse testcase.json")
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageExec,
		Env: []string{
			"TEST_CASE_DIR=" + filepath.Join("/code-grader/testcases", gr.TestCaseName),
		},
		OpenStdin: true,
		StdinOnce: true,
	}, &container.HostConfig{
		NetworkMode: "none",
		Resources: container.Resources{
			Memory: CompilationMemoryLimit,
		},
	}, nil, nil, "")

	if err != nil {
		return nil, errors.Wrap(InternalError, "cannot create container")
	}

	hjresp, err := cli.ContainerAttach(ctx, resp.ID, dockertypes.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})

	if err != nil {
		return nil, errors.Wrap(InternalError, "cannot attach container")
	}

	ctxdl, cancel := context.WithTimeout(ctx, time.Second*time.Duration(testCase.RuntimeOptions.RuntimeLimit))
	defer cancel()

	statusCh, errCh := cli.ContainerWait(ctxdl, resp.ID, container.WaitConditionNextExit)

	err = cli.ContainerStart(ctxdl, resp.ID, dockertypes.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(InternalError, "cannot start container")
	}

	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	hjresp.Conn.Write([]byte(gr.UserCode))
	hjresp.Conn.Close()
	stdcopy.StdCopy(outW, errW, hjresp.Conn)
	outW.Close()
	errW.Close()

	select {
	case status := <-statusCh:
		switch status.StatusCode {
		case 0:
			result, err := io.ReadAll(outR)
			if err != nil {
				return nil, errors.Wrap(InternalError, "error reading outR")
			}
			return result, nil
		case 1:
			result, err := io.ReadAll(errR)
			if err != nil {
				return nil, errors.Wrap(InternalError, "error reading errR")
			}
			return result, CompilationError
		case 2:
			return nil, InternalError
		default:
			return nil, InternalError
		}
	case err := <-errCh:
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, TimeLimitExceed
		}
		return nil, errors.Wrap(InternalError, "error waiting container")
	}

}
