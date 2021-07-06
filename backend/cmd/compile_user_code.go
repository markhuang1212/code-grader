package cmd

import (
	"context"
	"io"
	"os"
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
	Ok  bool
	Msg string
}

const CompilationMemoryLimit = 256 * 1024 * 1024
const CompilationTimeLimit = 10 * time.Second

// var ErrCompilationError = errors.New("compilation error")

// it is required that the docker image is built before the program runs
const imageCompile = "markhuang1212/code-grader/runtime-compile:latest"

// The function compiles user's code inside a docker container, and returns the
// executable on success
func CompileUserCode(ctx context.Context, gr types.GradeRequest, tmpDir string) (*CompileUserCodeResult, error) {

	result := &CompileUserCodeResult{}

	if !IsTestcase(gr.TestCaseName) {
		return nil, types.ErrNoTestCase
	}

	err := os.WriteFile(filepath.Join(tmpDir, "code.txt"), []byte(gr.UserCode), 0666)
	if err != nil {
		return nil, errors.Wrap(err, "cannot write code.txt")
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
	}, &container.HostConfig{
		Binds:       []string{tmpDir + ":/data"},
		NetworkMode: "none",
		Resources: container.Resources{
			Memory:     CompilationMemoryLimit,
			MemorySwap: CompilationMemoryLimit,
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
		err := cli.ContainerRemove(ctxdl, resp.ID, dockertypes.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			log.Panicln("cannot kill and remove container")
			panic(err)
		}
	}()

	select {
	case status := <-statusCh:
		switch status.StatusCode {
		case 0:
			result.Ok = true
			result.Msg = "compilation success"
			return result, nil
		case 1:
			out, err := cli.ContainerLogs(ctxdl, resp.ID, dockertypes.ContainerLogsOptions{
				ShowStderr: true,
				ShowStdout: true,
			})
			if err != nil {
				return nil, errors.Wrap(err, "error reading stdout")
			}
			errText, _ := io.ReadAll(out)
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
