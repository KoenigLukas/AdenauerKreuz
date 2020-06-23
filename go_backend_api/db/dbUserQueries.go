package db

import "log"

func RegisterUser(username string, email string, firstname string, lastname string, password string) int {
	_, err := Con.Query(`INSERT INTO users(username,email,first_name,last_name,password) values(?,?,?,?,?)`, username, email, firstname, lastname, password)
	if err != nil {
		return -1
	}
	var id int
	err = Con.QueryRow(`SELECT LAST_INSERT_ID()`).Scan(&id)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	return id;
}

func LoginUser(username string, password string) int{
	var id int
	err := Con.QueryRow(`SELECT userID FROM users WHERE username = ? and password = ?`,username,password).Scan(&id)
	if err != nil{
		log.Fatal(err)
		return -1
	}

	return id
}

func CheckUsernameAvailable(username string) (bool,error){
	var available bool
	err := Con.QueryRow(`SELECT EXISTS (SELECT * FROM users WHERE username = ?)`,username).Scan(&available)
	if err != nil{
		return false,err
	}
	return !available,err
}

func CheckEmailAvailable(email string) (bool,error){
	var available bool
	err := Con.QueryRow(`SELECT EXISTS (SELECT * FROM users WHERE email = ?)`,email).Scan(&available)
	if err != nil{
		return false,err
	}
	return !available,err
}

