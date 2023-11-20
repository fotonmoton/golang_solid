package o

import "fmt"

type Vehicle interface {
	AddPassenger(string)
}

type Car struct {
	passengers []string
}

func (c *Car) AddPassenger(p string) {
	c.passengers = append(c.passengers, p)
	fmt.Printf("Passenger %s has been boarded to the car!\n", p)
}

func (c *Car) BoardPassenger(p string) {
	c.passengers = append(c.passengers, p)
}

type Train struct {
	passengers []string
}

func (t *Train) AddPassenger(p string) {
	t.passengers = append(t.passengers, p)
	fmt.Printf("Passenger %s has been boarded to the train!\n", p)
}

func (t *Train) TakeASeat(p string) {
	t.passengers = append(t.passengers, p)
}

type Plane struct {
	passengers []string
}

func (plane *Plane) AddPassenger(p string) {
	plane.passengers = append(plane.passengers, p)
	fmt.Printf("Passenger %s has been boarded to the plain!\n", p)

}

func (plane *Plane) GoOnBoard(p string) {
	plane.passengers = append(plane.passengers, p)
}

func NotExtendableRide(vehicles ...Vehicle) {
	for _, vehicle := range vehicles {
		if car, ok := vehicle.(*Car); ok {
			car.BoardPassenger("Greg")
			fmt.Println("Passenger Greg has been boarded to the car!")
		}
		if train, ok := vehicle.(*Train); ok {
			train.TakeASeat("Greg")
			fmt.Println("Passenger Greg has been boarded to the train!")
		}
	}
}

func ExtendableRide(vehicles ...Vehicle) {
	for _, vehicle := range vehicles {
		vehicle.AddPassenger("Greg")
	}
}
