package mysql

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(user_id) from users where username = ?"
	//query, err := db.Query(sqlStr, username)
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}

	return count > 0, nil
}

// InsertUser 插入数据库
func InsertUser() {
	//
}
