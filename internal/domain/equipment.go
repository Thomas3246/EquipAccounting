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
	CPU          CPU
	GPU          GPU
	Motherboard  Motherboard
	RAM          int
	Storage      int
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

type EquipmentState struct {
	Id   int
	Name string
}
