package db

func RegisterUser(username string, email string, firstname string, lastname string, password string) int {
	_, err := Con.Query("INSERT INTO users(username,email,first_name,last_name,password) values($1,$2,$3,$4,$5)", username, email, firstname, lastname, password)
	if err != nil {
		return -1
	}
	var id int
	err = Con.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		return -1
	}
	return id;
}

func LoginUser(username string, password string) int{
	var id int
	err := Con.QueryRow("SELECT userid FROM users WHERE username = $1 and password = $2").Scan(&id)
	if err != nil{
		return -1
	}
	return id
}

func CheckUsernameAvailable(username string) (bool,error){
	var available bool
	err := Con.QueryRow("SELECT EXISTS (SELECT * FROM users WHERE username = $1)",username).Scan(&available)
	if err != nil{
		return false,err
	}
	return !available,err
}

func CheckEmailAvailable(email string) (bool,error){
	var available bool
	err := Con.QueryRow("SELECT EXISTS (SELECT * FROM users WHERE email = $1)",email).Scan(&available)
	if err != nil{
		return false,err
	}
	return !available,err
}

