package l

import "fmt"

type Employee interface {
	Name() string
}

type HR struct{}

func (HR) Name() string {
	return "HR Pamela"
}

func (h *HR) Hire(e Employee) {
	fmt.Println(e.Name(), "was hired!")
}

type Administrator struct{}

func (Administrator) Name() string {
	return "Administrator Bob"
}

type Manager struct{}

func (Manager) Name() string {
	return "Manager Jack"
}

type CEO struct {
	Manager
}
