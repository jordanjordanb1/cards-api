package decks

import "github.com/google/uuid"

type card struct {
	face  string
	suit  string
	value int
}

type Deck struct {
	deck_id   string
	shuffled  bool
	remaining int
	cards     []card
}

func NewDeck() *Deck {
	var cards []card
	uuid := uuid.New().String()

	// Gets all cards without map key
	for _, card := range deckOfCards {
		cards = append(cards, card)
	}

	deck := &Deck{deck_id: uuid, shuffled: false, remaining: 52, cards: cards}

	return deck
}

var deckOfCards = map[string]card{
	"KS": {
		face:  "king",
		suit:  "spades",
		value: 13,
	},
	"QS": {
		face:  "queen",
		suit:  "spades",
		value: 12,
	},
	"10S": {
		face:  "ten",
		suit:  "spades",
		value: 10,
	},
	"9S": {
		face:  "nine",
		suit:  "spades",
		value: 9,
	},
	"8S": {
		face:  "eight",
		suit:  "spades",
		value: 8,
	},
	"7S": {
		face:  "seven",
		suit:  "spades",
		value: 7,
	},
	"6S": {
		face:  "six",
		suit:  "spades",
		value: 6,
	},
	"5S": {
		face:  "five",
		suit:  "spades",
		value: 5,
	},
	"4S": {
		face:  "four",
		suit:  "spades",
		value: 4,
	},
	"3S": {
		face:  "three",
		suit:  "spades",
		value: 3,
	},
	"2S": {
		face:  "two",
		suit:  "spades",
		value: 2,
	},
	"AS": {
		face:  "ace",
		suit:  "spades",
		value: 1,
	},
	"KC": {
		face:  "king",
		suit:  "clubs",
		value: 13,
	},
	"QC": {
		face:  "queen",
		suit:  "clubs",
		value: 12,
	},
	"10C": {
		face:  "ten",
		suit:  "clubs",
		value: 10,
	},
	"9C": {
		face:  "nine",
		suit:  "clubs",
		value: 9,
	},
	"8C": {
		face:  "eight",
		suit:  "clubs",
		value: 8,
	},
	"7C": {
		face:  "seven",
		suit:  "clubs",
		value: 7,
	},
	"6C": {
		face:  "six",
		suit:  "clubs",
		value: 6,
	},
	"5C": {
		face:  "five",
		suit:  "clubs",
		value: 5,
	},
	"4C": {
		face:  "four",
		suit:  "clubs",
		value: 4,
	},
	"3C": {
		face:  "three",
		suit:  "clubs",
		value: 3,
	},
	"2C": {
		face:  "two",
		suit:  "clubs",
		value: 2,
	},
	"AC": {
		face:  "ace",
		suit:  "clubs",
		value: 1,
	},
	"KD": {
		face:  "king",
		suit:  "diamonds",
		value: 13,
	},
	"QD": {
		face:  "queen",
		suit:  "diamonds",
		value: 12,
	},
	"10D": {
		face:  "ten",
		suit:  "diamonds",
		value: 10,
	},
	"9D": {
		face:  "nine",
		suit:  "diamonds",
		value: 9,
	},
	"8D": {
		face:  "eight",
		suit:  "diamonds",
		value: 8,
	},
	"7D": {
		face:  "seven",
		suit:  "diamonds",
		value: 7,
	},
	"6D": {
		face:  "six",
		suit:  "diamonds",
		value: 6,
	},
	"5D": {
		face:  "five",
		suit:  "diamonds",
		value: 5,
	},
	"4D": {
		face:  "four",
		suit:  "diamonds",
		value: 4,
	},
	"3D": {
		face:  "three",
		suit:  "diamonds",
		value: 3,
	},
	"2D": {
		face:  "two",
		suit:  "diamonds",
		value: 2,
	},
	"AD": {
		face:  "ace",
		suit:  "diamonds",
		value: 1,
	},
	"KH": {
		face:  "king",
		suit:  "hearts",
		value: 13,
	},
	"QH": {
		face:  "queen",
		suit:  "hearts",
		value: 12,
	},
	"10H": {
		face:  "ten",
		suit:  "hearts",
		value: 10,
	},
	"9H": {
		face:  "nine",
		suit:  "hearts",
		value: 9,
	},
	"8H": {
		face:  "eight",
		suit:  "hearts",
		value: 8,
	},
	"7H": {
		face:  "seven",
		suit:  "hearts",
		value: 7,
	},
	"6H": {
		face:  "six",
		suit:  "hearts",
		value: 6,
	},
	"5H": {
		face:  "five",
		suit:  "hearts",
		value: 5,
	},
	"4H": {
		face:  "four",
		suit:  "hearts",
		value: 4,
	},
	"3H": {
		face:  "three",
		suit:  "hearts",
		value: 3,
	},
	"2H": {
		face:  "two",
		suit:  "hearts",
		value: 2,
	},
	"AH": {
		face:  "ace",
		suit:  "hearts",
		value: 1,
	},
}
