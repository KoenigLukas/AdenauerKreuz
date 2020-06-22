package handler

import (
	"encoding/json"
	"net/http"
)


func ListCreateHandler(w http.ResponseWriter, r *http.Request){

}

func ListGetAllHandler(w http.ResponseWriter, r *http.Request){

}

func ListGetHandler(w http.ResponseWriter, r *http.Request){

}

func ListEntryAddHandler(w http.ResponseWriter, r *http.Request){

}

func ListEntryEditHandler(w http.ResponseWriter, r *http.Request){

}

func ListEntryDeleteHandler(w http.ResponseWriter, r *http.Request){

}







type ListElement struct {
	ElementID   int    `json:"elementID"`
	Score       int    `json:"score"`
	Content     string `json:content`
	Connotation bool   `json:"connotation"`
}

type List struct {
	ListID       int            `json:"listID"`
	Name         string         `json:"name"`
	ListElements *[]ListElement `json:"elements"`
}

type AllLists struct{
	Count int     `json:"count"`
	Lists *[]List `json:"lists"`
}

type UpdateListElem struct{
	ListID int               `json:"listID"`
	ListElement *ListElement `json:"entry""`
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	var listElems []ListElement
	item := ListElement{1, 10, "schlechtes Wetter", false}
	item2 := ListElement{2, 15, "nicht zu warm", true}
	listElems = append(listElems, item, item2)

	list := List{1, "TestList", &listElems}

	var listElems2 []ListElement
	item3 := ListElement{1,5,"test",true}
	item4 := ListElement{2,10,"test2",false}
	listElems2 = append(listElems2, item3,item4)

	list2 := List{2,"TestList2",&listElems2}

	var lists []List
	lists = append(lists,list,list2)

	all := AllLists{2,&lists}

	//listupdate := UpdateListElem{2,&item}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(all)
}
