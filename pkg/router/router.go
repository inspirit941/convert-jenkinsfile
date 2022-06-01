package router

import (
	"github.com/gin-gonic/gin"
	"github.com/inspirit941/convert-jenkinsfile/docs"
	"github.com/inspirit941/convert-jenkinsfile/pkg/api"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(server *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := server.Group("/api/v1")
	{
		v1.POST("/upload", api.ConvertFile)
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Use(CORSMiddleware())
	return server
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://delightful-field-0835ff900.1.azurestaticapps.net")
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
