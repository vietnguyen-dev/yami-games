package main

import (
	"fmt"
)

func main() {
	var player1 string
	var player2 string

	fmt.Println("What is player 1 name?")
	fmt.Scanln(&player1)
	fmt.Println("What is player 2 name?")
	fmt.Scanln(&player2)

	fmt.Println("The players are: ", player1, "and ", player2)
	fmt.Println("GAME START")
}
