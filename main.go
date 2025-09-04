package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Card struct {
	code   string
	image  string
	images Images
	value  string
	suit   string
}

type Images struct {
	svg string
	png string
}

type Deck struct {
	success   string
	deck_id   string
	shuffed   bool
	remaining int
}

func main() {
	var player1 string
	var player2 string

	p1 := []Card{}
	p2 := []Card{}

	fmt.Println("What is player 1 name?")
	fmt.Scanln(&player1)
	fmt.Println("What is player 2 name?")
	fmt.Scanln(&player2)

	fmt.Println("The players are: ", player1, "and ", player2)
	fmt.Println("GAME START")

	response, err := http.Get("https://www.deckofcardsapi.com/api/deck/new/shuffle/?deck_count=1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	var deck Deck
	if err := json.NewDecoder(response.Body).Decode(&deck); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	api := fmt.Sprintf("https://www.deckofcardsapi.com/api/deck/%s/draw/?count=26", deck.deck_id)

	p1Cards, err := http.Get(api)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer p1Cards.Body.Close()
}
