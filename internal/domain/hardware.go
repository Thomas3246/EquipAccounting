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

type Unit interface {
	GetID() int
	GetName() string
}

type CPU struct {
	Id   int
	Name string
}

func (c CPU) GetID() int      { return c.Id }
func (c CPU) GetName() string { return c.Name }

type GPU struct {
	Id   int
	Name string
}

func (g GPU) GetID() int      { return g.Id }
func (g GPU) GetName() string { return g.Name }

type Motherboard struct {
	Id   int
	Name string
}

func (m Motherboard) GetID() int      { return m.Id }
func (m Motherboard) GetName() string { return m.Name }
