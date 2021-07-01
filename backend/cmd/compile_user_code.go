package cmd

import (
	"context"
	"io"
	//"log"
	"path/filepath"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/markhuang1212/code-grader/backend/types"
	"github.com/pkg/errors"
)

// Memory Limit of a G++ Process
const CompilationMemoryLimit = 100 * 1024 * 1024

// it is required that the docker image is built before the program runs
const imageCompile = "markhuang1212/code-grader/runtime-compile:latest"

// The function compiles user's code inside a docker container, and returns the
// executable on success
func CompileUserCode(ctx context.Context, gr types.GradeRequest) ([]byte, error) {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, errors.WithMessage(err, "cannot create docker client")
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageCompile,
		Env: []string{
			"TEST_CASE_DIR=" + filepath.Join("/code-grader/testcases", gr.TestCaseName),
			"CXX=g++",
			"CXXFLAGS=-std=c++11",
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
		return nil, errors.Wrap(ErrInternalError, "cannot create container")
	}

	attachInput, err := cli.ContainerAttach(ctx, resp.ID, dockertypes.ContainerAttachOptions{
		Stdin: true,
	})

	if err != nil {
		return nil, errors.Wrap(ErrInternalError, "cannot attach container input")
	}

	attachOutput, err := cli.ContainerAttach(ctx, resp.ID, dockertypes.ContainerAttachOptions{
		Stdout: true,
		Stderr: true,
	})

	if err != nil {
		return nil, errors.Wrap(ErrInternalError, "cannot attach container output")
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNextExit)

	err = cli.ContainerStart(ctx, resp.ID, dockertypes.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(ErrInternalError, "cannot start container")
	}

	defer func() {
		// err := cli.ContainerRemove(ctx, resp.ID, dockertypes.ContainerRemoveOptions{
		// 	Force: true,
		// })
		// if err != nil {
		// 	log.Fatal("cannot kill and remove container")
		// 	panic(err)
		// }
	}()

	attachInput.Conn.Write([]byte(gr.UserCode))
	attachInput.Close()

	outR, outW := io.Pipe()
	errR, errW := io.Pipe()
	stdcopy.StdCopy(outW, errW, attachOutput.Conn)
	outW.Close()
	errW.Close()
	attachOutput.Conn.Close()

	select {
	case status := <-statusCh:
		switch status.StatusCode {
		case 0:
			result, err := io.ReadAll(outR)
			if err != nil {
				return nil, errors.Wrap(ErrInternalError, "error reading outR")
			}
			return result, nil
		case 1:
			result, err := io.ReadAll(errR)
			if err != nil {
				return nil, errors.Wrap(ErrInternalError, "error reading errR")
			}
			return result, ErrCompilationError
		case 2:
			return nil, ErrInternalError
		default:
			return nil, ErrInternalError
		}
	case <-errCh:
		return nil, errors.Wrap(ErrInternalError, "error waiting container")
	}
}
