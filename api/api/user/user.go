package user

import (
	"github.com/gin-gonic/gin"
	"zg5/lmonth/api/conts"
	"zg5/lmonth/api/service"
	"zg5/lmonth/framework/http"
)

func Login(c *gin.Context) {
	var loginReq struct {
		Tel  string `json:"Tel"`
		Code string `json:"Code"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	if loginReq.Tel == "" {
		http.Re(c, conts.PRM_ERROR, "手机号不能为空", nil)
		return
	}
	if loginReq.Code == "" {
		http.Re(c, conts.PRM_ERROR, "验证码不能为空", nil)
		return
	}
	login, err := service.Login(c, loginReq.Tel, loginReq.Code)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	http.Re(c, conts.SUCCESS, "登陆成功", login)
}

func Register(c *gin.Context) {
	var loginReq struct {
		Tel      string `json:"Tel"`
		Password string `json:"Password"`
		Code     string `json:"Code"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
	}
	if loginReq.Password == "" || loginReq.Tel == "" || loginReq.Code == "" {
		http.Re(c, conts.PRM_ERROR, "请求参数有误", nil)
		return
	}
	user, err := service.RegisterUser(c, loginReq.Tel, loginReq.Password, loginReq.Code)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	http.Re(c, conts.SUCCESS, "注册成功", user)
}

func Message(c *gin.Context) {
	var loginReq struct {
		Tel string `json:"Tel"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
	}
	err := service.SendMessage(c, loginReq.Tel)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	http.Re(c, conts.SUCCESS, "发送成功", nil)
	return
}
