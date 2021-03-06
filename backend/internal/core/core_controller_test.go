package core_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/markhuang1212/code-grader/backend/internal/core"
	"github.com/markhuang1212/code-grader/backend/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCoreController(t *testing.T) {

	cc := core.NewCoreController(1)

	cc.GradeQueue <- types.GradeRequest{
		Id:           "test1",
		TestCaseName: "example-1",
		UserCode:     "int main() { cout << \"Hello\" << endl; }",
	}

	for {
		val, ok := cc.Cache.Get("test1")
		if ok == true {
			assert.Equal(t, types.GradeResultSuccess, val.Status)
			return
		}
		runtime.Gosched()
	}
}

func BenchmarkCoreController(b *testing.B) {
	threads := 20
	cc := core.NewCoreController(threads)
	for i := 0; i < b.N; i++ {
		cc.GradeQueue <- types.GradeRequest{
			Id:           "test" + fmt.Sprint(i),
			TestCaseName: "example-1",
			UserCode:     "int main() { cout << \"Hello\" << endl; }",
		}
	}
	for {
		if cc.Cache.Count() == b.N {
			return
		}
		runtime.Gosched()
	}
}
