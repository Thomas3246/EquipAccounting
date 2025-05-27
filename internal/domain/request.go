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
}
