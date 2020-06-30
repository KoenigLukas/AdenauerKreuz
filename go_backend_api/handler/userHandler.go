package handler

import (
	"encoding/json"
	"github.com/koeniglukas/db"
	"github.com/koeniglukas/storage"
	"net/http"
	"strconv"
)

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var req storage.ChangePwReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := storage.ErrorMsg{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	idstr := r.Header.Get("userID")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}


	err = db.ChangePassword(req.Old_password,idint,req.New_password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func ChangeEmailHandler(w http.ResponseWriter, r *http.Request) {
	var req storage.ChangeEmailReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errmsg := storage.ErrorMsg{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	idstr := r.Header.Get("userID")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}


	err = db.ChangePassword(req.Old_email,idint,req.New_email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.Header.Get("userID")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	err = db.DeleteUser(idint)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.WriteHeader(http.StatusOK)
	return

}
