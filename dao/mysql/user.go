package mysql

func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(user_id) from users where username = ?"
	//query, err := db.Query(sqlStr, username)
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}

	return count > 0, nil
}

func InsertUser() {
	//
}
