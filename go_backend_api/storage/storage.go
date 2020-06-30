package storage

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

type RegisterErr struct {
	Err_message    string `json:"err_message"`
	Email_check    bool   `json:"email_check"`
	Username_check bool   `json:"username_check"`
}

type TokenStore struct {
	Token string `json:"token"`
}

type AvailableCheck struct {
	Available bool `json:"available"`
}

type ChangePwReq struct {
	Old_password string `json:"old_password"`
	New_password string `json:"new_password"`
}

type ChangeEmailReq struct {
	Old_email string `json:"old_email"`
	New_email string `json:"new_email"`
}

type ListElemRq struct {
	ListID      int         `json:listID`
	ListElement ListElement `json:"entry"`
}

type ListElement struct {
	ElementID   int    `json:"elementID"`
	Score       int    `json:"score"`
	Content     string `json:content`
	Connotation bool   `json:"connotation"`
}

type List struct {
	ListID       int           `json:"listID"`
	Name         string        `json:"name"`
	ListElements []ListElement `json:"elements"`
}

type Lists struct {
	Count int    `json:"count"`
	Lists []List `json:"lists"`
}

type ListCreateReq struct {
	Name string `json:"name"`
}

type ListCreateRes struct {
	ListID int `json:"listID"`
}

type DelEntryReq struct {
	ListID int `json:"listID"`
	ElementID int `json:"elementID"`
}
