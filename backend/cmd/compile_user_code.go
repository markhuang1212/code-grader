package cmd

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"log"
	"path/filepath"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/markhuang1212/code-grader/backend/types"
	"github.com/pkg/errors"
)

type CompileUserCodeResult struct {
	Ok     bool
	Result []byte
	Msg    string
}

const CompilationMemoryLimit = 256 * 1024 * 1024
const CompilationTimeLimit = 10 * time.Second

// var ErrCompilationError = errors.New("compilation error")

// it is required that the docker image is built before the program runs
const imageCompile = "markhuang1212/code-grader/runtime-compile:latest"

// The function compiles user's code inside a docker container, and returns the
// executable on success
func CompileUserCode(ctx context.Context, gr types.GradeRequest) (*CompileUserCodeResult, error) {

	result := &CompileUserCodeResult{}

	if !IsTestcase(gr.TestCaseName) {
		return nil, types.ErrNoTestCase
	}

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, errors.WithMessage(err, "cannot create docker client")
	}

	ctxdl, cancel := context.WithTimeout(ctx, CompilationTimeLimit)
	defer cancel()

	resp, err := cli.ContainerCreate(ctxdl, &container.Config{
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
		return nil, errors.Wrap(err, "cannot create container")
	}

	hjresp, err := cli.ContainerAttach(ctxdl, resp.ID, dockertypes.ContainerAttachOptions{
		Stdin:  true,
		Stream: true,
	})

	if err != nil {
		return nil, errors.Wrap(err, "cannot attach container")
	}

	statusCh, errCh := cli.ContainerWait(ctxdl, resp.ID, container.WaitConditionNextExit)

	err = cli.ContainerStart(ctxdl, resp.ID, dockertypes.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "cannot start container")
	}

	defer func() {
		err := cli.ContainerRemove(ctxdl, resp.ID, dockertypes.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			log.Panicln("cannot kill and remove container")
			panic(err)
		}
	}()

	userCodeLength := len(gr.UserCode)
	fmt.Fprintln(hjresp.Conn, strconv.Itoa(userCodeLength))
	hjresp.Conn.Write([]byte(gr.UserCode))
	hjresp.Close()

	select {
	case status := <-statusCh:
		switch status.StatusCode {
		case 0:
			stdout, err := cli.ContainerLogs(ctxdl, resp.ID, dockertypes.ContainerLogsOptions{
				ShowStdout: true,
				ShowStderr: false,
			})
			if err != nil {
				return nil, errors.Wrap(err, "error reading stdout")
			}
			program, _ := io.ReadAll(stdout)
			result.Ok = true
			result.Result = program
			return result, nil
		case 1:
			stderr, err := cli.ContainerLogs(ctxdl, resp.ID, dockertypes.ContainerLogsOptions{
				ShowStderr: true,
				ShowStdout: false,
			})
			if err != nil {
				return nil, errors.Wrap(err, "error reading stdout")
			}
			errText, _ := io.ReadAll(stderr)
			result.Msg = string(errText)
			result.Ok = false
			return result, nil
		case 2:
			return nil, types.ErrInternal
		default:
			return nil, types.ErrInternal
		}
	case err := <-errCh:
		return nil, errors.Wrap(err, "error waiting container")
	}
}
