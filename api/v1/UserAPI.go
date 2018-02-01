package v1

import (
	"github.com/gin-gonic/gin"
	"ginWebDemo/api"
	"ginWebDemo/api/util"
	"ginWebDemo/api/database"
)

type PasswordLogin struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
}

func LoginUser(g gin.IRoutes) {
	g.POST("/user/login", func(context *gin.Context) {
		var l PasswordLogin
		err := context.BindJSON(&l)
		api := &api.HttpError{}
		if err != nil {
			util.Convert(err)
			api.Message = "缺少必要参数"
			api.SendError(context)
			return
		}
		if ! util.ValidatePhone(l.Phone) {
			api.Message = "电话号码格式不正确"
			api.SendError(context)
			return
		}
		if !database.UserValidate(l.Phone) {
			api.Message = "用户不存在"
			api.SendError(context)
		}
		user, ok := database.UserLogin(l.Phone, l.Password)
		if ok {
			api.Message = "登录成功"
			api.Data = user
			api.SendData(context)
		} else {
			api.Message = "手机号码或密码不正确，登录失败"
			api.SendError(context)
		}
	})
}

func RegisterUser(g gin.IRoutes) {
	g.POST("/user/register", func(context *gin.Context) {
		var l PasswordLogin
		err := context.BindJSON(&l)
		api := &api.HttpError{}
		if err != nil {
			util.Convert(err)
			api.Message = "缺少必要参数"
			api.SendError(context)
			return
		}
		if ! util.ValidatePhone(l.Phone) {
			api.Message = "电话号码格式不正确"
			api.SendError(context)
			return
		}
		if database.UserValidate(l.Phone) {
			api.Message = "用户已存在"
			api.SendError(context)
		}
		ok := database.RegisterUser(l.Phone, l.Password)
		if ok {
			api.Message = "注册成功"
			api.SendData(context)
		} else {
			api.Message = "注册用户失败"
			api.SendError(context)
		}
	})
}
