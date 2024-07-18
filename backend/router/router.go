package router

import (
	"backend/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(testController *controller.TestController) *gin.Engine {
	router := gin.Default()
	// add swagger
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")

	testRouter := baseRouter.Group("/test")
	testRouter.GET("", testController.FindAll)
	testRouter.GET("/:testId", testController.FindById)
	testRouter.POST("", testController.Create)
	testRouter.PATCH("/:testId", testController.Update)
	testRouter.DELETE("/:testId", testController.Delete)

	return router
}