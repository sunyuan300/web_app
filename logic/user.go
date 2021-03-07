package logic

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// 存放业务的逻辑代码

func SignUp(p *models.ParamSign) (err error) {
	// 1.判断用户存不存在
	var exist bool
	exist, err = mysql.CheckUserExist(p.UserName)
	if err != nil {
		// 数据库查询出错
		return err
	}
	if exist {
		// 用户已存在错误
		return errors.New("用户已存在")
	}
	// 2.生成UID
	UserID := snowflake.GenID()

	// 3.密码加密

	// 构造一个User实例
	user := &models.User{
		UserID:   UserID,
		UserName: p.UserName,
		Password: p.Password,
	}
	// 4.保存进数据库
	mysql.InsertUser(user)
}
