package domain

type Hardware struct {
	Id                   int
	EquipmentDirectoryId int
	CPU                  CPU
	GPU                  GPU
	Motherboard          Motherboard
	RAM                  int
	Storage              int
}

type CPU struct {
	Id   int
	Name string
}

type GPU struct {
	Id   int
	Name string
}

type Motherboard struct {
	Id   int
	Name string
}
