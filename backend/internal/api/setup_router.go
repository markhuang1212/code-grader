package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markhuang1212/code-grader/backend/internal/core"
	"github.com/markhuang1212/code-grader/backend/internal/grader"
	"github.com/markhuang1212/code-grader/backend/internal/types"
	"github.com/markhuang1212/code-grader/backend/internal/util"
)

const ApiPrefix = "/api/v1"

type GradeResultResponse struct {
	Ready  bool
	Result types.GradeResult
}

func SetupRouter(cc *core.CoreController) *gin.Engine {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/api", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://app.swaggerhub.com/apis-docs/markhuang1212/CodeGraderCore")
	})

	authorized := r.Group(ApiPrefix)

	authorized.POST("/grade", func(c *gin.Context) {

		gr := types.GradeRequest{}

		err := c.BindJSON(&gr)
		if err != nil {
			return
		}

		gr.Id = util.RandomHex(10)

		cc.GradeQueue <- gr

		c.Header("Location", ApiPrefix+"/result/"+gr.Id)
		c.Status(202)

	})

	authorized.GET("/result/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		result, ok := cc.Cache.Get(id)
		if !ok {
			c.JSON(200, GradeResultResponse{
				Ready: false,
			})
		} else {
			c.JSON(200, GradeResultResponse{
				Ready:  true,
				Result: result,
			})
		}
	})

	authorized.GET("/testcases", func(c *gin.Context) {
		result := grader.LoadTestcases()
		c.JSON(200, result)
	})

	return r
}
