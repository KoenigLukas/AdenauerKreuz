package handler

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/koeniglukas/config"
	"github.com/koeniglukas/db"
	"github.com/koeniglukas/storage"
	"net/http"
)


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var login storage.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		errmsg := storage.ErrorMsg{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	id := db.LoginUser(login.Username, login.Password)
	if id == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{"database error"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	var token storage.TokenStore

	token.Token, err = createToken(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{"token generation failed"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
	return

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var register storage.Register
	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		errmsg := storage.ErrorMsg{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	if !(register.Username_check && register.Email_check) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		regErrMsg := storage.RegisterErr{"Invalid Checks", register.Email_check, register.Username_check}
		json.NewEncoder(w).Encode(regErrMsg)
		return
	}

	id := db.RegisterUser(register.Username, register.Email, register.FirstName, register.LastName, register.Password)
	if id == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{"database error"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	var token storage.TokenStore

	token.Token, err = createToken(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{"token generation failed"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
	return
}

func CheckUsernameHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	available, err := db.CheckUsernameAvailable(pathParams["username"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}
	ret := storage.AvailableCheck{available}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
	return
}

func CheckEmailHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	available, err := db.CheckEmailAvailable(pathParams["email"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}
	ret := storage.AvailableCheck{available}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
	return
}

type Claims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

func createToken(userid int) (string, error) {
	var err error
	claims := Claims{UserID: userid}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(config.Get("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
