package userService

import (
	"gin-api/utils/db"
	"gin-api/utils/log"
	"golang.org/x/crypto/bcrypt"
	"gin-api/models"
	"fmt"
	"gin-api/utils/redis"
	"gin-api/utils/e"
	"github.com/asaskevich/govalidator"
	"gin-api/utils/jwt"
	"github.com/mojocn/base64Captcha"
	"time"
	"regexp"
	"math/rand"
)

const MobileRegex  = "^1[0-9]{10}$"
const CaptchaExpire  = 5 * 60

const JwtExpireTime  = 30 * time.Minute
const KeyExpireTime = 7 * 24 * time.Hour

//登录
func Login(loginForm LoginForm) (string, string, int) {
	var user *models.User
	var errCode int
	if govalidator.IsNull(loginForm.Code) {
		user, errCode = loginByPassword(loginForm)
	} else {
		user, errCode = loginByCode(loginForm)
	}
	var token string
	var key string
	if errCode == e.SUCCESS{
		var err error
		token, err = jwt.GenerateToken(int(user.ID), JwtExpireTime)
		key, err = jwt.GenerateToken(int(user.ID), KeyExpireTime)
		if err != nil {
			log.Warnf("[Login] create token fail: %v", err)
			errCode = e.USER_INTER_ERROR
		}
	}
	return token, key, errCode
}

//验证码登录
func loginByCode(loginForm LoginForm) (*models.User, int) {
	user := models.User{}
	b := validateCode(loginForm.Mobile, loginForm.Code)
	if !b {
		return nil, e.USER_LOGIN_ERROR
	} else {
		q := db.ORM.Where("mobile = ?", loginForm.Mobile).First(&user)
		if err := q.Error; err != nil{
			log.Warnf("query fail: %v", err)
			return nil, e.USER_LOGIN_ERROR
		}
	}
	return &user, e.SUCCESS
}

//密码登录
func loginByPassword(loginForm LoginForm) (*models.User, int)  {
	user := models.User{}

	log.Warn(loginForm.Mobile)

	q := db.ORM.Table("user").Where("mobile = ?", loginForm.Mobile).First(&user)
	if err := q.Error; err != nil{
		log.Warnf("query fail: %v", err)
		return nil, e.USER_LOGIN_ERROR
	}
	b := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginForm.Password))
	if b != nil {
		log.Warnf("login fail: %v", b)
		return nil, e.USER_LOGIN_ERROR
	} else {
		return &user, e.SUCCESS
	}
}

//创建新用户
func CreateUser(registerForm RegisterForm) (string, string, int){
	user := models.User {
		Mobile: registerForm.Mobile,
		Password: registerForm.Password,
	}

	b := validateCode(registerForm.Mobile, registerForm.Code)
	if !b {
		log.Warnf("vcode invalidate")
		return "","",e.USER_VCODE_ERROR
	}

	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		log.Warnf("password hash fail: %v", err)
		return "","", e.USER_INTER_ERROR
	}
	user.Password = string(hpass)
	err = db.ORM.Create(&user).Error
	if err != nil {
		log.Warnf("user create fail: %v", err)
		return "","", e.USER_INTER_ERROR
	}
	var token string
	var key string
	token, err = jwt.GenerateToken(int(user.ID), JwtExpireTime)
	key, err = jwt.GenerateToken(int(user.ID), KeyExpireTime)
	if err != nil {
		log.Warnf("[Login] create token fail: %v", err)
		return "","", e.USER_INTER_ERROR
	}
	return token, key, e.SUCCESS
}

func validateCode(mobile, code string) bool{
	key := fmt.Sprintf("users:vcode:%s", mobile)

	c, err := redis.GetString(key)
	if err != nil {
		log.Warnf("[validateCode] error:%v", err)
	}
	b := c == code
	if b {
		redis.Delete(key)
	}
	return b

}

/*
 * 发送验证码
 */
func SendVCode(mobile, captcha string) int  {
	key := fmt.Sprintf("users:capatcha:%s", mobile)
	idKey, err := redis.GetString(key)

	if err != nil {
		log.Warnf("[SendSMSCode] get redis key:%s error:%v", key, err)
		return e.USER_INVALID_PARAMS
	}
	if b := verfiyCaptcha(idKey, captcha); !b {
		log.Warn("[SendSMSCode] verify captcha error")
		redis.Delete(key)
		return e.USER_CAPTCHA_ERROR
	}

	vcode := createVCode()
	vCodeKey := fmt.Sprintf("users:vcode:%s", mobile)
	_, err = redis.Set(vCodeKey, vcode, "EX", CaptchaExpire)
	if err != nil {
		log.Warnf("[SendSMSCode] set redis error:%v", err)
		return e.USER_INTER_ERROR
	}
	return e.SUCCESS
}

/*
 * 获取图片验证码
 */
func GetCapatcha(mobile string) (string, int) {
	idKeyD, base64StringD := createCapatcha()

	key := fmt.Sprintf("users:capatcha:%s", mobile)

	_, err := redis.Set(key, idKeyD, "EX", CaptchaExpire)
	if err != nil {
		log.Warnf("[GetCaptcha] set redis error:%v", err)
		return "", e.USER_INTER_ERROR
	}
	return base64StringD, e.SUCCESS
}

/*
 * 创建图片验证码
 */
func createCapatcha() (idKeyD, base64StringD string) {
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	base64StringD = base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return
}

/*
 * 验证图片验证码
 */
func verfiyCaptcha(idkey,verifyValue string) bool{
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	return verifyResult
}

/*
 * 创建六位数字验证码
 */
func createVCode() (code string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code = fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return
}

func validatorMobile(mobile string) bool {
	b, _ := regexp.MatchString(MobileRegex, mobile)
	return b
}