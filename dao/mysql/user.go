package mysql

import "web_app/models"

// 把每一步数据库操作封装成函数,待logic层根据业务需求调用

func CheckUserExist(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

func InsertUser(user *models.User) {
	// 执行SQL语句入库
}
