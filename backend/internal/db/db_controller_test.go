package db_test

import (
	"context"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/markhuang1212/code-grader/backend/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	redisBin, err := exec.LookPath("redis-server")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(redisBin)
	cmd.Dir = "/data"
	cmd.Start()
	time.Sleep(time.Second)

	m.Run()

	cmd.Process.Signal(os.Interrupt)
	cmd.Process.Wait()

}

func TestDbControllerInit(t *testing.T) {
	ctx := context.Background()
	c := db.NewDbController("localhost:6379")
	c.LoadVersion(ctx)
	time.Sleep(time.Second)
	val, err := c.Rdb.Get(ctx, "version").Result()
	assert.Nil(t, err)
	assert.Equal(t, "v0.1.0", val)
}
