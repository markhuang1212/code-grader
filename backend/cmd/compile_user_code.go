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

// Memory Limit of a G++ Process
const CompilationMemoryLimit = 100 * 1024 * 1024

// it is required that the docker image is built before the program runs
const imageName = "markhuang1212/code-grader/runtime-compile"

var InternalError = errors.New("internal error")
var CompilationFailure = errors.New("compilation error")

// The function compiles user's code inside a docker container, and returns the
// executable on success
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
		AttachStderr: true,
		OpenStdin:    true,
	}, &container.HostConfig{
		NetworkMode: "none",
		Resources:   container.Resources{
			// Memory: CompilationMemoryLimit,
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

	statucCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNextExit)

	err = cli.ContainerStart(ctx, resp.ID, dockertypes.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(InternalError, "cannot start container")
	}

	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	hjresp.Conn.Write([]byte(gr.UserCode))
	hjresp.Conn.Close()
	stdcopy.StdCopy(outW, errW, hjresp.Conn)

	select {
	case status := <-statucCh:
		if status.StatusCode == 0 {
			out, err := io.ReadAll(outR)
			if err != nil {
				return nil, errors.Wrap(InternalError, "cannot read stdout")
			}
			return out, err
		} else {
			out, err := io.ReadAll(errR)
			if err != nil {
				return nil, errors.Wrap(InternalError, "cannot read stderr")
			}
			return out, CompilationFailure
		}
	case <-errCh:
		return nil, errors.Wrap(InternalError, "error waiting container")
	}
}
