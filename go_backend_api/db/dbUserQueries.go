package db

func RegisterUser(username string, email string, firstname string, lastname string, password string) int {
	rows, err := Con.Query("INSERT INTO users(username,email,first_name,last_name,password) values(?,?,?,?,?)", username, email, firstname, lastname, password)
	if err != nil {
		return -1
	}
	rows.Close()
	var id int
	err = Con.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		return -1
	}
	return id
}

func LoginUser(username string, password string) int {
	var id int
	err := Con.QueryRow("SELECT userid FROM users WHERE username = ? and password = ?",username,password).Scan(&id)
	if err != nil {
		return -1
	}
	return id
}

func CheckUsernameAvailable(username string) (bool, error) {
	var available bool
	err := Con.QueryRow("SELECT EXISTS (SELECT * FROM users WHERE username = ?)", username).Scan(&available)
	if err != nil {
		return false, err
	}
	return !available, err
}

func CheckEmailAvailable(email string) (bool, error) {
	var available bool
	err := Con.QueryRow("SELECT EXISTS (SELECT * FROM users WHERE email = ?)", email).Scan(&available)
	if err != nil {
		return false, err
	}
	return !available, err
}

func ChangePassword(old_password string, userid int, new_password string) error {
	rows, err := Con.Query("ALTER TABLE users SET password = ? WHERE password = ? and userid = ?", new_password, old_password, userid)
	rows.Close()
	return err
}

func ChangeEmail(old_email string, userid int, new_email string) error {
	rows, err := Con.Query("ALTER TABLE users SET email = ? WHERE email = ? and userid = ?", new_email, old_email, userid)
	rows.Close()
	return err
}

func DeleteUser(userid int) error {
	rows, err := Con.Query("DELETE FROM users WHERE userID = ?", userid)
	rows.Close()
	return err
}
