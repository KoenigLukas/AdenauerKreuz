package handler

import (
	"github.com/koeniglukas/config"
	"github.com/koeniglukas/db"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username_check bool   `json:"username_check"`
	Email_check    bool   `json:"email_check"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Password       string `json:"password"`
}

type ErrorMsg struct {
	Err_message string `json:"err_message"`
}

type RegisterErr struct{
	Err_message string `json:"err_message"`
	Email_check    bool   `json:"email_check"`
	Username_check bool   `json:"username_check"`
}

type TokenStore struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		errmsg := ErrorMsg{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	id := db.LoginUser(login.Username, login.Password)
	if id == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := ErrorMsg{"database error"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	var token TokenStore

	token.Token, err = createToken(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := ErrorMsg{"token generation failed"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
	return

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var register Register
	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		errmsg := ErrorMsg{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	if !(register.Username_check && register.Email_check){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		regErrMsg := RegisterErr{"Invalid Checks",register.Email_check,register.Username_check}
		json.NewEncoder(w).Encode(regErrMsg)
		return
	}

	id := db.RegisterUser(register.Username,register.Email,register.FirstName,register.LastName,register.Password)
	if id == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := ErrorMsg{"database error"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	var token TokenStore

	token.Token, err = createToken(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := ErrorMsg{"token generation failed"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
	return
}

func CheckUsernameHandler(w http.ResponseWriter, r *http.Request) {

}

func CheckEmailHandler(w http.ResponseWriter, r *http.Request) {

}

func createToken(userid int) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.Get("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

