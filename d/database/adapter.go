package database

type DatabaseAdapter struct {
	Connection *DatabaseConnection
}

func (d *DatabaseAdapter) GetPossibleMoves() []string {
	moves := []string{}

	for _, move := range d.Connection.PossibleMovesV2 {
		moves = append(moves, move.Name)
	}

	return moves
}
