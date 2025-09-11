package deck

import (
	"math/rand"
	"time"
)

type Card struct {
	Value string
	Suit  string
	Color string
}

type Deck struct {
	Cards []Card
}

var (
	suits  = [4]string{"SPADE", "CLOVER", "DIAMOND", "HEART"}
	values = [14]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING", "ACE", "JOKER"}
)

func manualShuffle(slice []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range slice {
		ri := r.Intn(len(slice))
		slice[i], slice[ri] = slice[ri], slice[i]
	}
	return slice
}

func (d Deck) Init() {
	d.Cards = []Card{}
	for _, suit := range suits {
		for _, value := range values {
			var color string
			if suit == "SPADE" || suit == "CLOVER" {
				color = "BLACK"
			} else {
				color = "RED"
			}
			newCard := Card{
				Value: value,
				Suit:  suit,
				Color: color,
			}
			d.Cards = append(d.Cards, newCard)
		}
	}
	d.Cards = manualShuffle(d.Cards)
}

func (d Deck) Draw() Card {
	deckLength := len(d.Cards)
	newLength := deckLength - 1
	cardToDraw := d.Cards[deckLength]
	d.Cards = d.Cards[0:newLength]
	return cardToDraw
}
