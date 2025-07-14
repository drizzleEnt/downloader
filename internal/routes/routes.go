package routes

import (
	"downloader/internal/api"

	"github.com/gin-gonic/gin"
)

func InitRoutes(controller api.Controller) *gin.Engine {
	r := gin.New()

	api := r.Group("/files")
	api.POST("/download", func(ctx *gin.Context) {
		controller.Download(ctx.Request.Context(), ctx.Writer, ctx.Request)
	})
	api.POST("/batch/download", func(ctx *gin.Context) {

	})

	return r
}
