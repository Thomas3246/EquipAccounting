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
}

type RequestType struct {
	Id   int
	Name string
}

type RequestResult struct {
	Id   int
	Name string
}
