package models

// 定义请求参数的结构体

type ParamSign struct {
	UserName   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
}
