package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpError struct {
	Message string
}

func (h *HttpError) SendError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": -2,
		"msg":  h,
	})
}
