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
	return server
}
