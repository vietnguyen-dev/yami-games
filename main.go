package main

import (
	"fmt"

	"github.com/vietnguyen-dev/yami-games/deck"
)

type Game struct {
	Player1name     string
	Player2name     string
	player1position int
	player2position int
	CardsChosen     []deck.Card
}

func (g Game) setPlayer1Name(playerName string) {
	g.Player1name = playerName
}

func (g Game) setPlayer2Name(playerName string) {
	g.Player2name = playerName
}

func main() {
	d := deck.Init{}
	fmt.Println("GAME START")
}
