package v1

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	"net/http"
	"fmt"
	"regexp"
	"ginWebDemo/api"
)

type Phone struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
}

func SetupRounter() *gin.Engine {
	r := gin.Default()
	g := r.Group("/miapp/api/v1")
	g.POST("/getSmsCode", func(c *gin.Context) {
		var phone Phone
		err := c.BindJSON(&phone)
		api := &api.HttpError{}
		if ( err != nil) {
			fmt.Println(err)
			//log.Fatal(err)
			api.Message = "缺少必要参数"
			api.SendError(c)
			return
		}
		m, e := regexp.MatchString("1[3|5|7|8|9][0-9]{9}", phone.Phone)
		if ( !m || e != nil) {
			api.Message = "电话号码格式不正确"
			api.SendError(c)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":    "获取验证码成功",
			"status": 200,
			"data": gin.H{
				"code": generateSmsCode(),
			},
		})
	})
	return r
}

func generateSmsCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprint(rnd.Int31n(1000000))
}
