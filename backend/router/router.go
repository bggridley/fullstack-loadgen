package router

import (
	"backend/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
  
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
  
		c.Next()
	}
  }

func NewRouter(testController *controller.TestController) *gin.Engine {
	router := gin.Default()
	// add swagger
	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(CORSMiddleware())
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