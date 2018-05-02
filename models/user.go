package models


type User struct {
	Model
	Mobile		string	`json:"mobile"`
	NickName	string	`json:"nick_name"`
	Password 	string 	`json:"password"`
}


