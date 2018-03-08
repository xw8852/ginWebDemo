package v1

import "github.com/gin-gonic/gin"

func Handle(g *gin.Engine){
	group := g.Group("/api/v1")
	SendMsgCode(group)
	LoginUser(group)
	RegisterUser(group)
	WechatLogin(group)
}