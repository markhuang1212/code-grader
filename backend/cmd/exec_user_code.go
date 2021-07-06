package cmd

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"

	"github.com/markhuang1212/code-grader/backend/types"
)

const imageExec = "markhuang1212/code-grader/runtime-exec:latest"

type ExecUserCodeResult struct {
	Ok           bool
	MemoryExceed bool
	TimeExceed   bool
	Msg          string
}

func ExecUserCode(ctx context.Context, gr types.GradeRequest, tmpDir string) (*ExecUserCodeResult, error) {

	result := &ExecUserCodeResult{}

	if !IsTestcase(gr.TestCaseName) {
		return nil, types.ErrNoTestCase
	}

	testcaseConfJson, err := os.ReadFile(filepath.Join(GetAppRoot(), "testcases", gr.TestCaseName, "testcase.json"))
	if err != nil {
		return nil, errors.Wrap(err, "cannot read testcase.json")
	}

	testcaseConf := types.TestCaseOptions{}
	err = json.Unmarshal(testcaseConfJson, &testcaseConf)
	if err != nil {
		return nil, errors.Wrap(err, "cannot parse testcase.json")
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, errors.WithMessage(err, "cannot create docker client")
	}

	ctxdl, cancel := context.WithTimeout(ctx, time.Duration(testcaseConf.RuntimeOptions.RuntimeLimit)*time.Second)
	defer cancel()

	resp, err := cli.ContainerCreate(ctxdl, &container.Config{
		Image: imageExec,
		Env: []string{
			"TEST_CASE_DIR=" + filepath.Join("/code-grader/testcases", gr.TestCaseName),
		},
	}, &container.HostConfig{
		NetworkMode: "none",
		Binds:       []string{tmpDir + ":/data"},
		Resources: container.Resources{
			Memory:     int64(testcaseConf.RuntimeOptions.MemoryLimit) * 1024 * 1024,
			MemorySwap: int64(testcaseConf.RuntimeOptions.MemoryLimit) * 1024 * 1024,
		},
	}, nil, nil, "")
	if err != nil {
		return nil, errors.Wrap(err, "cannot create container")
	}

	statusCh, errCh := cli.ContainerWait(ctxdl, resp.ID, container.WaitConditionNextExit)

	err = cli.ContainerStart(ctxdl, resp.ID, dockertypes.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "cannot start container")
	}

	defer func() {
		err := cli.ContainerRemove(ctx, resp.ID, dockertypes.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			log.Panicln("cannot kill and remove container")
		}
	}()

	select {
	case status := <-statusCh:
		switch status.StatusCode {
		case 0:
			result.Ok = true
			result.Msg = "correct answer"
			return result, nil
		case 1, 2:
			out, err := cli.ContainerLogs(ctxdl, resp.ID, dockertypes.ContainerLogsOptions{
				ShowStderr: true,
				ShowStdout: true,
			})
			if err != nil {
				return nil, errors.Wrap(err, "error reading stdout")
			}
			text, _ := io.ReadAll(out)
			result.Msg = string(text)
			result.Ok = false
			return result, nil
		case 3:
			return nil, types.ErrInternal
		case 137:
			result.Ok = false
			result.Msg = "memory limit exceed"
			result.MemoryExceed = true
			return result, nil
		default:
			return nil, types.ErrInternal
		}
	case <-errCh:
		result.Ok = false
		result.Msg = "time limit exceed"
		result.TimeExceed = true
		return result, nil
	}

}
