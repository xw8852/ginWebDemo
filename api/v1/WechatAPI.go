package v1

import (
	"github.com/gin-gonic/gin"
	"ginWebDemo/api"
	"ginWebDemo/api/util"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"ginWebDemo/api/database"
	"time"
	"math/rand"
	"fmt"
)

type WechatCode struct {
	Code string `form:"Code" json:"code" binding:"required"`
}
type WeChatInfo struct {
	Openid      string  `json:"openid"`
	Session_key string  `json:"session_key"`
	Unionid     string  `json:"unionid"`
	Errcode     int  `json:"errcode"`
	Errmsg      string  `json:"errmsg"`
}

/**
* 微信
* code 换取 session_key
 */
func WechatLogin(g gin.IRoutes) {
	g.POST("/wechat/login", func(c *gin.Context) {
		var code WechatCode
		err := c.BindJSON(&code)
		api := &api.HttpError{}
		if err != nil {
			util.Convert(err)
			api.Message = "缺少必要参数"
			api.SendError(c)
			return
		}
		response, _ := http.Get("https://api.weixin.qq.com/sns/jscode2session?" +
			"appid=wxade5302ee9685fc8" +
			"&secret=1037d10b8d1da6374606f511ee15cd25" +
			"&js_code=" + code.Code + "&grant_type=authorization_code")
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
		if response.StatusCode == 200 {
			wechatinfo := WeChatInfo{}
			json.Unmarshal(body, &wechatinfo)
			if (wechatinfo.Openid != "") {
				user, ok := database.UserGetByWeChat(wechatinfo.Openid)
				if (ok) {
					api.Data = user
					api.Message = "用户登录成功"
					return
				}
				logonName := ""
				for {
					t := time.Time{}
					index := database.UserCount()
					index = index + t.Year() - t.Minute() - t.Second() - rand.Intn(100)
					logonName = "wx" + string(index)
					ok := database.RegisterUserByName(logonName, "123456")
					if ok {
						break
					}
				}
				user, _ = database.UserLoginByName(logonName, "123456")
				database.InsertWeChatRelation(user.Id, wechatinfo.Openid)
				api.Data = user
				api.Message = "用户登录成功"
				return
			} else {
				api.Message = wechatinfo.Errmsg
				api.SendError(c)
				return
			}
		} else {
			api.Message = "用户微信登录失败"
			api.SendError(c)
			return
		}
	})
}
