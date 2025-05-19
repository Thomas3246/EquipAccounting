package domain

type User struct {
	Id       int
	Login    string
	Password string // HASH ONLY
	Role     string
}
