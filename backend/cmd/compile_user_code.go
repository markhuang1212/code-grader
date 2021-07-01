package cmd

import (
	"context"
	"io"
	"strconv"

	//"log"
	"path/filepath"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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
	}, &container.HostConfig{
		NetworkMode: "none",
		Resources:   container.Resources{
			// Memory: CompilationMemoryLimit,
		},
	}, nil, nil, "")

	if err != nil {
		return nil, errors.Wrap(ErrInternalError, "cannot create container")
	}

	hjresp, err := cli.ContainerAttach(ctx, resp.ID, dockertypes.ContainerAttachOptions{
		Stdin:  true,
		Stream: true,
	})

	if err != nil {
		return nil, errors.Wrap(ErrInternalError, "cannot attach container")
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

	userCodeLength := len(gr.UserCode)
	hjresp.Conn.Write([]byte(strconv.Itoa(userCodeLength) + "\n"))
	hjresp.Conn.Write([]byte(gr.UserCode))
	hjresp.Close()

	if err != nil {
		return nil, errors.Wrap(ErrInternalError, "cannot close attached session (output)")
	}

	select {
	case status := <-statusCh:
		switch status.StatusCode {
		case 0:
			stdout, err := cli.ContainerLogs(ctx, resp.ID, dockertypes.ContainerLogsOptions{
				ShowStdout: true,
				ShowStderr: false,
			})
			if err != nil {
				return nil, errors.Wrap(ErrInternalError, "error reading stdout")
			}
			result, _ := io.ReadAll(stdout)
			return result, nil
		case 1:
			stderr, err := cli.ContainerLogs(ctx, resp.ID, dockertypes.ContainerLogsOptions{
				ShowStderr: true,
				ShowStdout: false,
			})
			if err != nil {
				return nil, errors.Wrap(ErrInternalError, "error reading stdout")
			}
			result, _ := io.ReadAll(stderr)
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
