package db

import "github.com/koeniglukas/storage"




func CreateList(userid int, name string) int {
	var listid int
	_, err := Con.Query("INSERT INTO lists(userid,name) VALUES(?,?)", userid, name)
	if err != nil {
		return -1
	}
	err = Con.QueryRow("SELECT LAST_INSERT_ID();").Scan(&listid)
	if err != nil {
		return -1
	}
	return listid
}

func GetAllLists(userID int) (storage.Lists,error){
	var lists storage.Lists

	rows,err := Con.Query("SELECT listID,name from lists where userid = ?",userID)
	if err != nil{
		return lists,err
	}
	defer rows.Close()
	for rows.Next() {
		var list storage.List
		rows.Scan(&list.ListID,&list.Name)
		elemRow, err := Con.Query("SELECT elementID,score,content,connotation From list_element Where listid = ?",list.ListID)
		if err != nil {
			return lists,err
		}
		for elemRow.Next(){
			var element storage.ListElement
			elemRow.Scan(&element.ElementID,&element.Score,&element.Content,&element.Connotation)
			list.ListElements = append(list.ListElements,element)
		}
		elemRow.Close()
		lists.Lists = append(lists.Lists, list)
		lists.Count += 1
	}
	return lists,nil
}

func GetList(userID int, listID int) (storage.List,error) {
	var list storage.List
	list.ListID = listID
	err := Con.QueryRow("SELECT name from lists where listID = ?",listID).Scan(&list.Name)
	if err != nil {
		return list,err
	}

	rows,err := Con.Query("SELECT elementID,score,content,connotation From list_element Where listid = ?",listID)
	if err != nil {
		return list,err
	}
	defer rows.Close()
	for rows.Next(){
		var element storage.ListElement
		rows.Scan(&element.ElementID,&element.Score,&element.Content,&element.Connotation)
		list.ListElements = append(list.ListElements,element)
	}
	return list,nil
}

func AddListEntry(elem storage.ListElement,listid int) (int,error){
	row, err := Con.Query("INSERT INTO list_element(listID,content,score,orientation) VALUES(?,?,?,?)",listid,elem.Content,elem.Score,elem.Connotation)
	defer row.Close()
	if err != nil {
		return -1,err
	}

	var id int
	err = Con.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		return -1,err
	}
	return id,nil
}

func EditListEntry(elem storage.ListElement,listID int) error{
	tx,err := Con.Begin()
	if err != nil{
		return err
	}

	_,err = tx.Exec("DELETE FROM list_element WHERE elementID = ? and listID = ?",elem.ElementID,listID)
	if err != nil{
		tx.Rollback()
		return err
	}
	_,err = tx.Exec("INSERT INTO list_element(elementid,listID,content,score,orientation) VALUES(?,?,?,?,?)",elem.ElementID,listID,elem.Content,elem.Score,elem.Connotation)
	if err != nil{
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

func DeleteListEntry(elementID int, listID int) error{
	row,err := Con.Query("DELETE FROM list_element WHERE elementID = ? and listID = ?",elementID,listID)
	defer row.Close()
	return err
}