package userController

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"github.com/asaskevich/govalidator"
	"gin-api/utils/log"
	"gin-api/controllers/response"
	"gin-api/utils/e"
	"gin-api/service/userService"
)

const MobileRegex  = "^1[0-9]{10}$"




// @Summary 注册
// @Description 新用户注册
// @Tags User
// @Produce json
// @Param mobile query string true "Mobile"
// @Param password query string false "Password"
// @Param code query string false "Code"
// @Success 200 {string} json "{"succeed":true,"data":{},"errorCode":0,"errorMessage":""}"
// @Failure 200 {string} json "{"succeed":false,"data":{},"errorCode":100002,"errorMessage":"请求参数错误"}"
// @Router /users/register [post]
func Register(c *gin.Context)  {
	var registerForm userService.RegisterForm
	if err := c.ShouldBind(&registerForm); err != nil {
		log.Warnf("[Register] error: %v", err)
		response.E(c, e.USER_INVALID_PARAMS)
		return
	}
	token, key, errCode := userService.CreateUser(registerForm)
	if errCode == e.SUCCESS {
		response.J(c, gin.H{
			"token": token,
			"key": key,
		})
	} else {
		response.E(c, errCode)
	}
}

// @Summary 登录
// @Description 密码登录，验证码登录
// @Tags User
// @Produce json
// @Param mobile query string true "Mobile"
// @Param password query string false "Password"
// @Param code query string false "Code"
// @Success 200 {string} json "{"succeed":true,"data":{},"errorCode":0,"errorMessage":""}"
// @Failure 200 {string} json "{"succeed":false,"data":{},"errorCode":100002,"errorMessage":"请求参数错误"}"
// @Router /users/login [post]
func Login(c *gin.Context)  {
	var loginForm userService.LoginForm
	if err := c.ShouldBind(&loginForm); err != nil {
		response.E(c, e.USER_INVALID_PARAMS)
		return
	}
	if govalidator.IsNull(loginForm.Code) && govalidator.IsNull(loginForm.Password){
		response.E(c, e.USER_INVALID_PARAMS)
		return
	}
	token, key, errCode := userService.Login(loginForm)
	if errCode == e.SUCCESS {
		response.J(c, gin.H{
			"token": token,
			"key": key,
		})
	} else {
		response.E(c, errCode)
	}
}


// @Summary 发送短信验证码
// @Description 发送短信验证码
// @Tags User
// @Produce json
// @Param mobile path string true "Mobile"
// @Param captcha query string false "Captcha"
// @Success 200 {string} json "{"succeed":true,"data":{},"errorCode":0,"errorMessage":""}"
// @Failure 200 {string} json "{"succeed":false,"data":{},"errorCode":100002,"errorMessage":"请求参数错误"}"
// @Router /users/vcode/{mobile} [get]
func SendSMSCode(c *gin.Context)  {
	mobile:= c.Param("mobile")
	captcha := c.Query("captcha")

	if b, _ := regexp.MatchString(MobileRegex, mobile); !b {
		log.Warnf("[GetCaptcha] illegal mobile:%s", mobile)
		response.E(c, e.USER_INVALID_PARAMS)
		return
	}


	if govalidator.IsNull(captcha) {
		log.Warn("[SendSMSCode] captcha is null")
		response.E(c, e.USER_INVALID_PARAMS)
		return
	}
	errCode := userService.SendVCode(mobile, captcha)
	if errCode == e.SUCCESS{
		response.J(c, nil)
	} else {
		response.E(c, errCode)
	}
}

// @Summary 获取图片验证码
// @Description 获取图片验证码
// @Tags User
// @Produce json
// @Param mobile path string true "Mobile"
// @Success 200 {string} json "{"succeed":true,"data":{},"errorCode":0,"errorMessage":""}"
// @Failure 200 {string} json "{"succeed":false,"data":{},"errorCode":100002,"errorMessage":"请求参数错误"}"
// @Router /users/captcha/{mobile} [get]
func GetCaptcha(c *gin.Context)  {
	mobile := c.Param("mobile")
	if b, _ := regexp.MatchString(MobileRegex, mobile); !b {
		log.Warnf("[GetCaptcha] illegal mobile:%s", mobile)
		response.E(c, e.USER_INVALID_PARAMS)
		return
	}
	base64StringD, errCode := userService.GetCapatcha(mobile)
	if errCode == e.SUCCESS{
		response.J(c, gin.H{
			"captcha": base64StringD,
		})
	} else {
		response.E(c, errCode)
	}
}







