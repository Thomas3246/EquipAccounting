package domain

type Department struct {
	Id       int
	Name     string
	Division int
}

type DepartmentView struct {
	Id           int
	Name         string
	DivisionName string
}
