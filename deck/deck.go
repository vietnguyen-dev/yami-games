package main

type Card struct {
	Value string
	Suit  string
}

type Deck struct {
	Cards []Card
}

var (
	suits  = [4]string{"SPADE", "CLOVER", "DIAMOND", "HEART"}
	values = [14]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "JACK", "QUEEN", "KING", "ACE", "JOKER"}
)

func (d Deck) initDeck() {
	d.Cards = []Card{}
	for _, suit := range suits {
		for _, value := range values {
			newCard := Card{
				Value: value,
				Suit:  suit,
			}
			d.Cards = append(d.Cards, newCard)
		}
	}
}

func (d Deck) shuffleDeck() {
}
