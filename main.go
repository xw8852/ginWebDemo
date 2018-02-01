package main

import (
	"ginWebDemo/api/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	v1.Handle(g)
	g.Run("0.0.0.0:8080")
}
