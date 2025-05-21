package domain

type Request struct {
	Id          int
	WorkStation int
	Type        int
	Description string
	Author      int
	Status      int
	CreatedAt   string
}
