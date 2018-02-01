package v1

import (
	"github.com/gin-gonic/gin"
	"ginWebDemo/api"
	"fmt"
	"net/http"
	"time"
	"math/rand"
	"ginWebDemo/api/database"
	"ginWebDemo/api/util"
)

type Phone struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
}

/**
发送手机验证码
 */
func SendMsgCode(g gin.IRoutes) {
	g.POST("/common/getSmsCode", func(c *gin.Context) {
		var phone Phone
		err := c.BindJSON(&phone)
		api := &api.HttpError{}
		if err != nil {
			util.Convert(err)
			api.Message = "缺少必要参数"
			api.SendError(c)
			return
		}
		if ! util.ValidatePhone(phone.Phone) {
			api.Message = "电话号码格式不正确"
			api.SendError(c)
			return
		}
		code := generateSmsCode()
		database.InsertSmsCode(phone.Phone, string(code))
		c.JSON(http.StatusOK, gin.H{
			"msg":    "获取验证码成功",
			"status": 200,
			"data": gin.H{
				"code": code,
			},
		})
	})
}
func generateSmsCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprint(rnd.Int31n(1000000))
}
