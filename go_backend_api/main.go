package main

import (
	"encoding/json"
	"fmt"
)

type ListElement struct {
	ElementID   int  `json:"elementID"`
	Score       int  `json:"score"`
	Connotation bool `json:"connotation"`
}

type List struct {
	ListID       int          `json:"listID"`
	Name         string          `json:"name"`
	ListElements *[]ListElement `json:"elements"`
}

func main(){
	var listElems []ListElement
	item := ListElement{1,10,false}
	item2 := ListElement{2,15,true}
	listElems = append(listElems, item,item2)

	list := List{1,"TestList",&listElems}


	data,err := json.Marshal(list)
	if(err != nil){
		panic(err)
	}
	dataString := string(data)
	fmt.Println(dataString)
}