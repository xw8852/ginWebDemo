package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpError struct {
	Message string
	Data    interface{}
}

func (h *HttpError) SendError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": -2,
		"msg":  h.Message,
	})
}
func (h *HttpError) SendData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  h.Message,
		"data": h.Data,
	})
}
