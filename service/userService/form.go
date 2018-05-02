package userService

type LoginForm struct {
	Mobile 		string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password 	string `form:"password" json:"password"`
	Code 		string `form:"code" json:"code"`
}

type RegisterForm struct {
	Mobile		string	`form:"mobile" json:"mobile" binding:"required,mobile"`
	Password    string	`form:"password" json:"password" binding:"required,password"`
	Code		string	`form:"code" json:"code" binding:"required"`
}


type SendVCodeForm struct {
	Mobile 		string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Captcha     string `form:"captcha" json:"captcha"`
}