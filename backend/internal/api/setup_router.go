package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markhuang1212/code-grader/backend/internal/core"
	"github.com/markhuang1212/code-grader/backend/internal/types"
	"github.com/markhuang1212/code-grader/backend/internal/util"
)

func SetupRouter(cc *core.CoreController) *gin.Engine {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	authorized := r.Group("/")

	authorized.POST("/grade", func(c *gin.Context) {

		gr := types.GradeRequest{}

		err := c.BindJSON(&gr)
		if err != nil {
			return
		}

		gr.Id = util.RandomHex(10)

		cc.GradeQueue <- gr

		c.Header("Location", "/api/result/"+gr.Id)
		c.Status(202)

	})

	return r
}
