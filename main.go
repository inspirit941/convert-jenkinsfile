package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inspirit941/convert-jenkinsfile/pkg/router"
)

func main() {
	server := gin.Default()
	// router 세팅
	server = router.InitRouter(server)

	server.Run(":8000")
}
