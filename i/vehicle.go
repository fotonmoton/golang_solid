package i

import "fmt"

type Vehicle interface {
	Ride() int     // returns current speed
	PlayMusic(int) // pass volume
}

type FormulaOne struct{}

func (f FormulaOne) Ride() int {
	return 300
}

func (FormulaOne) PlayMusic(int) {
	fmt.Println("No music!")
}

type Taxi struct{}

func (t Taxi) Ride() int {
	return 150
}

func (Taxi) PlayMusic(int) {
	fmt.Println("Playing some music!")
}

type PickupWithLoudspeakers struct{}

func (PickupWithLoudspeakers) Ride() int {
	return 50
}

func (PickupWithLoudspeakers) PlayMusic(int) {
	fmt.Println("Playing loud music from pickup truck!")
}

type Ridable interface {
	Ride() int // returns current speed

}

type Playable interface {
	PlayMusic(int) // pass volume
}

type MovingWithMusic interface {
	Ridable
	Playable
}
