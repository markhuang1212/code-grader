package cmd

import (
	"context"
	"io"
	"path/filepath"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/markhuang1212/code-grader/types"
	"github.com/pkg/errors"
)

const CompilationMemoryLimit = 50 * 1024 * 1024
const imageName = "markhuang1212/code-grader/runtime-compile"

var InternalError = errors.New("internal error")
var CompilationFailure = errors.New("compilation error")

func CompileUserCode(ctx context.Context, gr types.GradeRequest) ([]byte, error) {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, errors.WithMessage(err, "cannot create docker client")
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Env: []string{
			"TEST_CASE_DIR=" + filepath.Join("/code-grader/testcases", gr.TestCaseName),
		},
		AttachStdin:  true,
		AttachStdout: true,
	}, &container.HostConfig{
		NetworkMode: "none",
		Resources: container.Resources{
			Memory: CompilationMemoryLimit,
		},
	}, nil, nil, "")

	if err != nil {
		return nil, errors.Wrap(err, "cannot create container")
	}

	cli.ContainerStart(ctx, resp.ID, dockertypes.ContainerStartOptions{})

	hjresp, err := cli.ContainerAttach(ctx, resp.ID, dockertypes.ContainerAttachOptions{
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "cannot attach container")
	}

	outR, outW := io.Pipe()
	errR, errW := io.Pipe()

	stdcopy.StdCopy(outW, errW, hjresp.Reader)
	hjresp.Conn.Write([]byte(gr.UserCode))

	statucCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNextExit)
	select {
	case status := <-statucCh:
		if status.StatusCode == 0 {
			out, err := io.ReadAll(outR)
			if err != nil {
				return nil, errors.Wrap(err, "cannot read stdout")
			}
			return out, err
		} else {
			out, err := io.ReadAll(errR)
			if err != nil {
				return nil, errors.Wrap(err, "cannot read stderr")
			}
			return out, errors.New("compilation failed")
		}
	case err := <-errCh:
		return nil, err
	}
}
