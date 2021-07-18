package core

import (
	"context"

	"github.com/markhuang1212/code-grader/backend/internal/grader"
	"github.com/markhuang1212/code-grader/backend/internal/types"
)

type CoreController struct {
	Cache      *GradeResultCache
	GradeQueue chan types.GradeRequest
}

// A go routine that handles incoming grade
// requests endlessly
func (cc *CoreController) handleRequest() {
	for {
		request, ok := <-cc.GradeQueue

		if !ok {
			break
		}

		ctx := context.Background()
		result, err := grader.GradeUserCode(ctx, request)

		if err != nil {
			cc.Cache.Add(request.Id, types.GradeResult{
				Status: types.GradeResultInternalError,
				Msg:    "internal error",
			})
		} else {
			cc.Cache.Add(request.Id, *result)
		}
	}
}

func NewCoreController(concurrent int) *CoreController {

	ret := CoreController{
		Cache:      NewGradeResultCache(),
		GradeQueue: make(chan types.GradeRequest, 1024),
	}

	for i := 0; i < concurrent; i++ {
		go ret.handleRequest()
	}

	return &ret
}
