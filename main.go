package main

import (
	"ginWebDemo/api/v1"
	//"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	v1.SetupRounter().Run("8080")
}
