package i

import "fmt"

func NotSegregated() {
	taxi := Taxi{}
	f1 := FormulaOne{}

	var garage []Vehicle = []Vehicle{taxi, f1}

	fmt.Println(garage)

}

func Segregated() {
	taxi := Taxi{}
	f1 := FormulaOne{}
	pickup := PickupWithLoudspeakers{}

	var garage []Ridable = []Ridable{taxi, f1}

	fmt.Println("this vehicles are moved to the garage: ", garage)

	var carnivalVehicles []MovingWithMusic = []MovingWithMusic{taxi, pickup}

	fmt.Println("this vehicles are riding with music: ", carnivalVehicles)

	for _, v := range carnivalVehicles {
		v.PlayMusic(100)
	}
}
