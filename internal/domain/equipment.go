package domain

type Equipment struct {
	Id           int
	InvNum       string
	PurchDate    string
	RegDate      string
	DecomDate    string
	DirectoryId  int
	DepartmentId int
	StatusId     int
}

type EquipmentView struct {
	Id         int
	InvNum     string
	PurchDate  string
	RegDate    string
	DecomDate  string
	Directory  string
	Department string
	Status     string
}
