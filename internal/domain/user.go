package domain

type User struct {
	Id           int
	Name         string
	Login        string
	Password     string // HASH ONLY
	IsAdmin      int
	DepartmentId int
}

type ViewUser struct {
	Id         int
	Name       string
	Login      string
	IsAdmin    int
	Department string
}
