package db

type ListElement struct {
	ElementID   int  `json:"elementID"`
	Score       int  `json:"score"`
	Connotation bool `json:"connotation"`
}

type List struct {
	ListID       int            `json:"listID"`
	Name         string         `json:"name"`
	ListElements *[]ListElement `json:"elements"`
}

type MultLists struct {
	Count int     `json:"count"`
	Lists *[]List `json:lists`
}


func CreateList(userid int, name string) int {
	var listid int
	_, err := Con.Query("INSERT INTO lists(userid,name) VALUES($1,$2)", userid, name)
	if err != nil {
		return -1
	}
	err = Con.QueryRow("SELECT LAST_INSERT_ID();").Scan(&listid)
	if err != nil {
		return -1
	}
	return listid
}

//func GetAllLists(userid int) (MultLists,error) {
//	var lists MultLists
//
//	rows,err := Con.Query("SELECT listID,name FROM lists WHERE userID = $1",userid)
//	if err != nil{
//		return lists,err
//	}
//	for rows.Next(){
//		var list List
//		err = rows.Scan(&list.ListID,&list.Name)
//		if err != nil{
//			return lists,err
//		}
//
//
//	}
//
//}
