package l

import "fmt"

type Vehicle interface {
	Model() string
}

type Car struct {
	model string
}

func (c Car) Model() string {
	return c.model
}

type SportCar struct {
	Car
}

func (c SportCar) Model() int {
	return 1
}

type Garage struct {
	Vehicles []Vehicle
}

func (g *Garage) AddVehicle(v Vehicle) {
	g.Vehicles = append(g.Vehicles, v)
}

func (g *Garage) Show() {
	for _, vehicle := range g.Vehicles {
		fmt.Println(vehicle.Model(), "is in garage!")
	}
}
