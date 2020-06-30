package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/koeniglukas/db"
	"github.com/koeniglukas/storage"
	"net/http"
	"strconv"
)

func ListCreateHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.Header.Get("userID")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	var listCreateReq storage.ListCreateReq
	err = json.NewDecoder(r.Body).Decode(&listCreateReq)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}
	var listCreateRes storage.ListCreateRes
	listCreateRes.ListID = db.CreateList(idint, listCreateReq.Name)
	if listCreateRes.ListID == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{"Error creating List"}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listCreateRes)
	return

}

func ListGetAllHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.Header.Get("userID")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	lists,err := db.GetAllLists(idint)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
	return

}

func ListGetHandler(w http.ResponseWriter, r *http.Request) {
	idstr := r.Header.Get("userID")
	idint, err := strconv.Atoi(idstr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}
	pathParams := mux.Vars(r)
	listid, err := strconv.Atoi(pathParams["id"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	list,err := db.GetList(idint,listid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&list)
	return
}

func ListEntryAddHandler(w http.ResponseWriter, r *http.Request) {

	var addelement storage.ListElemRq
	err := json.NewDecoder(r.Body).Decode(&addelement)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	eid,err := db.AddListEntry(addelement.ListElement,addelement.ListID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}
	addelement.ListElement.ElementID = eid

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&addelement.ListElement)
	return
}

func ListEntryEditHandler(w http.ResponseWriter, r *http.Request) {
	var editelement storage.ListElemRq
	err := json.NewDecoder(r.Body).Decode(&editelement)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	err = db.EditListEntry(editelement.ListElement,editelement.ListID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&editelement.ListElement)
	return
}

func ListEntryDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var delElem storage.DelEntryReq
	err := json.NewDecoder(r.Body).Decode(&delElem)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		errmsg := storage.ErrorMsg{err.Error()}
		json.NewEncoder(w).Encode(errmsg)
		return
	}
	err = db.DeleteListEntry(delElem.ElementID,delElem.ListID)
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


func TestHandler(w http.ResponseWriter, r *http.Request) {

	var listElems []storage.ListElement
	item := storage.ListElement{1, 10, "schlechtes Wetter", false}
	item2 := storage.ListElement{2, 15, "nicht zu warm", true}
	listElems = append(listElems, item, item2)

	list := storage.List{1, "TestList", listElems}

	var listElems2 []storage.ListElement
	item3 := storage.ListElement{1, 5, "test", true}
	item4 := storage.ListElement{2, 10, "test2", false}
	listElems2 = append(listElems2, item3, item4)

	list2 := storage.List{2, "TestList2", listElems2}

	var lists []storage.List
	lists = append(lists, list, list2)

	all := storage.Lists{2, lists}

	//listupdate := UpdateListElem{2,&item}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all)
}
