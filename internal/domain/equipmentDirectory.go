package domain

type EquipmentDirectory struct {
	Id          int
	Name        string
	ReleaseYear string
	TypeId      int
}

type EquipmentDirectoryView struct {
	Id          int
	Name        string
	ReleaseYear string
	Type        string
}

type EquipmentDirectoryType struct {
	Id   int
	Name string
}
