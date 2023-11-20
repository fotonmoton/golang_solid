package d

import "solid/d/database"

func NotInverted() {
	db := database.DatabaseConnection{
		[]string{"Place a home on A1", "Take money from bank"},
		[]database.Move{{1, "Place a home on A1"}, {2, "Take money from bank"}},
	}
	monopoly := BadMonopoly{db}
	monopoly.StartGame()
}

func Inverted() {
	db := database.DatabaseConnection{
		[]string{"Place a home on A1", "Take money from bank"},
		[]database.Move{{1, "Place a home on A1"}, {2, "Take money from bank"}},
	}
	monopoly := InvertedMonopoly{&db}
	monopoly.StartGame()
}
