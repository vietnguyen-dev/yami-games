package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Card struct {
	Code   string `json:"code"`
	Image  string `json:"image"`
	Images Images `json:"Images"`
	Value  string `json:"Value"`
	Suit   string `json:"Suit"`
}

type Images struct {
	Svg string `json:"svg"`
	Png string `json:"png"`
}

type Deck struct {
	Success   bool   `json:"Suit"`
	DeckId    string `json:"deck_id"`
	Shuffed   bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type Draw struct {
	Success   bool   `json:"suit"`
	DeckId    string `json:"deck_id"`
	Cards     []Card `json:"cards"`
	Remaining int    `json:"remaining"`
}

func main() {
	var player1 string
	var player2 string

	fmt.Println("What is player 1 name?")
	fmt.Scanln(&player1)
	fmt.Println("What is player 2 name?")
	fmt.Scanln(&player2)

	fmt.Println("The players are: ", player1, "and ", player2)
	fmt.Println("GAME START")

	// getting all the cards
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

	api := fmt.Sprintf("https://www.deckofcardsapi.com/api/deck/%s/draw/?count=26", deck.DeckId)

	var p1Deck Draw
	p1Cards, err := http.Get(api)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer p1Cards.Body.Close()

	if err := json.NewDecoder(p1Cards.Body).Decode(&p1Deck); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var p2Deck Draw
	p2Cards, err := http.Get(api)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(2)
	}
	defer p2Cards.Body.Close()

	if err := json.NewDecoder(p2Cards.Body).Decode(&p2Deck); err != nil {
		fmt.Print(err.Error())
		os.Exit(2)
	}

	// game logic starts here graphics library?
}
