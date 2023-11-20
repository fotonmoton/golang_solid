package l

func NotSubstitutable() {
	g := Garage{}

	g.AddVehicle(Car{"Skoda Octavia"})
	// g.AddVehicle(any(SportCar{Car: Car{"McLaren F1"}}).(Vehicle))
	g.Show()
}

func Substitutable() {
	hr := HR{}
	hr.Hire(Administrator{})
	hr.Hire(Manager{})
	hr.Hire(HR{})
	hr.Hire(CEO{})
}
