package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
