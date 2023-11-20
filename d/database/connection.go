package database

type Move struct {
	Wight int
	Name  string
}

type DatabaseConnection struct {
	PossibleMoves   []string
	PossibleMovesV2 []Move
}

func (d *DatabaseConnection) GetPossibleMoves() []string {
	return d.PossibleMoves
}

func (d *DatabaseConnection) GetPossibleMovesV2() []Move {
	return d.PossibleMovesV2
}
