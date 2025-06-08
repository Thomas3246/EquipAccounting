package domain

type Document struct {
	Id        int
	RequestId int
	Type      int
	File      []byte
	AddDate   string
	UserId    int
	Name      string
}

type DocumentType struct {
	Id   int
	Name string
}

type DocumentView struct {
	Id        int
	RequestId int
	Type      string
	AddDate   string
	UserLogin string
	Name      string
}
