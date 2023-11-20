package d

import (
	"fmt"
	"solid/d/database"
)

type BadMonopoly struct {
	databaseConnection database.DatabaseConnection
}

func (bm *BadMonopoly) StartGame() {
	for _, move := range bm.databaseConnection.GetPossibleMoves() {
		bm.ProcessNextMove(move)
	}
}

func (bm *BadMonopoly) ProcessNextMove(move string) {
	fmt.Println("Next move is:", move)
}

type MoveFetcher interface {
	GetPossibleMoves() []string
}

type InvertedMonopoly struct {
	moveFetcher MoveFetcher
}

func (im *InvertedMonopoly) StartGame() {
	for _, move := range im.moveFetcher.GetPossibleMoves() {
		im.ProcessNextMove(move)
	}
}

func (im *InvertedMonopoly) ProcessNextMove(move string) {
	fmt.Println("Next move is:", move)
}
