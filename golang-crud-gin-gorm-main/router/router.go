package router

import (
	"golang-crud-gin/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PUT("/:id", tagsController.Update)
	tagsRouter.GET("/:id", tagsController.FindById)
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/filter", tagsController.Filter)
	return router
}
