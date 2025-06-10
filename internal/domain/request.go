package domain

type Request struct {
	Id          int
	Type        int
	Description string
	Author      int
	Status      int
	CreatedAt   string
	ClosedAt    string
	Equipment   int
	Result      int
	ResultDescr string
}

type RequestView struct {
	Id          int
	Type        string
	Description string
	Author      string
	Status      string
	CreatedAt   string
	ClosedAt    string
	Equipment   string
	Result      string
	ResultDescr string
}

type RequestType struct {
	Id   int
	Name string
}

type RequestResult struct {
	Id   int
	Name string
}

type RequestReport struct {
	RequestId   int
	TypeId      int
	Description string
	AdminName   string
	CreatedAt   string
	Equipment   EquipmentView
	ResultId    int
	ResultDescr string
	ReportDate  string
}
