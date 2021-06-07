package decks

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/jordanjordanb1/cards-api/cache"
)

type card struct {
	Face  string `json:"face"`
	Suit  string `json:"suit"`
	Value int    `json:"value"`
}

type Deck struct {
	Id        string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
	Cards     []card `json:"cards,omitempty"`
}

// Creates a new deck of cards with no custom selection
func NewDeck(shuffle bool) Deck {
	var cards []card
	uuid := uuid.New().String()
	shuffled := false

	// Gets all cards without map key
	for _, card := range deckOfCards {
		cards = append(cards, card)
	}

	if shuffle {
		// Shuffle cards
		cards = ShuffleDeck(cards)
		shuffled = true
	}

	deck := Deck{Id: uuid, Shuffled: shuffled, Remaining: 52, Cards: cards}

	return deck
}

// Creates a new deck based on selection passed as an argument
func NewCustomDeck(shuffle bool, cardSelection []string) Deck {
	var cards []card
	uuid := uuid.New().String()
	shuffled := false

	// Loops through array of passed card types
	for _, cardHash := range cardSelection {
		foundCard, found := deckOfCards[cardHash]

		// Checks if card type is valid
		if found {
			cards = append(cards, foundCard)
		}
	}

	if shuffle {
		// Shuffle cards
		cards = ShuffleDeck(cards)
		shuffled = true
	}

	deck := Deck{Id: uuid, Shuffled: shuffled, Remaining: len(cards), Cards: cards}

	return deck
}

// Saves deck in the in-memory cache
func (deck *Deck) Save() *Deck {
	cache.Cache.Set(deck.Id, deck, 1)

	return deck
}

// Get deck from cache
func GetDeckById(id string) (interface{}, error) {
	deck, found := cache.Cache.Get(id)

	// If deck not found, return error
	if !found {
		return nil, errors.New("No deck found with ID of " + id)
	}

	return deck, nil
}

// Shuffles deck
func ShuffleDeck(cards []card) []card {
	// Required to seed randomizer
	rand.Seed(time.Now().UnixNano())

	// Shuffles cards
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}

var deckOfCards = map[string]card{
	"KS": {
		Face:  "king",
		Suit:  "spades",
		Value: 13,
	},
	"QS": {
		Face:  "queen",
		Suit:  "spades",
		Value: 12,
	},
	"10S": {
		Face:  "ten",
		Suit:  "spades",
		Value: 10,
	},
	"9S": {
		Face:  "nine",
		Suit:  "spades",
		Value: 9,
	},
	"8S": {
		Face:  "eight",
		Suit:  "spades",
		Value: 8,
	},
	"7S": {
		Face:  "seven",
		Suit:  "spades",
		Value: 7,
	},
	"6S": {
		Face:  "six",
		Suit:  "spades",
		Value: 6,
	},
	"5S": {
		Face:  "five",
		Suit:  "spades",
		Value: 5,
	},
	"4S": {
		Face:  "four",
		Suit:  "spades",
		Value: 4,
	},
	"3S": {
		Face:  "three",
		Suit:  "spades",
		Value: 3,
	},
	"2S": {
		Face:  "two",
		Suit:  "spades",
		Value: 2,
	},
	"AS": {
		Face:  "ace",
		Suit:  "spades",
		Value: 1,
	},
	"KC": {
		Face:  "king",
		Suit:  "clubs",
		Value: 13,
	},
	"QC": {
		Face:  "queen",
		Suit:  "clubs",
		Value: 12,
	},
	"10C": {
		Face:  "ten",
		Suit:  "clubs",
		Value: 10,
	},
	"9C": {
		Face:  "nine",
		Suit:  "clubs",
		Value: 9,
	},
	"8C": {
		Face:  "eight",
		Suit:  "clubs",
		Value: 8,
	},
	"7C": {
		Face:  "seven",
		Suit:  "clubs",
		Value: 7,
	},
	"6C": {
		Face:  "six",
		Suit:  "clubs",
		Value: 6,
	},
	"5C": {
		Face:  "five",
		Suit:  "clubs",
		Value: 5,
	},
	"4C": {
		Face:  "four",
		Suit:  "clubs",
		Value: 4,
	},
	"3C": {
		Face:  "three",
		Suit:  "clubs",
		Value: 3,
	},
	"2C": {
		Face:  "two",
		Suit:  "clubs",
		Value: 2,
	},
	"AC": {
		Face:  "ace",
		Suit:  "clubs",
		Value: 1,
	},
	"KD": {
		Face:  "king",
		Suit:  "diamonds",
		Value: 13,
	},
	"QD": {
		Face:  "queen",
		Suit:  "diamonds",
		Value: 12,
	},
	"10D": {
		Face:  "ten",
		Suit:  "diamonds",
		Value: 10,
	},
	"9D": {
		Face:  "nine",
		Suit:  "diamonds",
		Value: 9,
	},
	"8D": {
		Face:  "eight",
		Suit:  "diamonds",
		Value: 8,
	},
	"7D": {
		Face:  "seven",
		Suit:  "diamonds",
		Value: 7,
	},
	"6D": {
		Face:  "six",
		Suit:  "diamonds",
		Value: 6,
	},
	"5D": {
		Face:  "five",
		Suit:  "diamonds",
		Value: 5,
	},
	"4D": {
		Face:  "four",
		Suit:  "diamonds",
		Value: 4,
	},
	"3D": {
		Face:  "three",
		Suit:  "diamonds",
		Value: 3,
	},
	"2D": {
		Face:  "two",
		Suit:  "diamonds",
		Value: 2,
	},
	"AD": {
		Face:  "ace",
		Suit:  "diamonds",
		Value: 1,
	},
	"KH": {
		Face:  "king",
		Suit:  "hearts",
		Value: 13,
	},
	"QH": {
		Face:  "queen",
		Suit:  "hearts",
		Value: 12,
	},
	"10H": {
		Face:  "ten",
		Suit:  "hearts",
		Value: 10,
	},
	"9H": {
		Face:  "nine",
		Suit:  "hearts",
		Value: 9,
	},
	"8H": {
		Face:  "eight",
		Suit:  "hearts",
		Value: 8,
	},
	"7H": {
		Face:  "seven",
		Suit:  "hearts",
		Value: 7,
	},
	"6H": {
		Face:  "six",
		Suit:  "hearts",
		Value: 6,
	},
	"5H": {
		Face:  "five",
		Suit:  "hearts",
		Value: 5,
	},
	"4H": {
		Face:  "four",
		Suit:  "hearts",
		Value: 4,
	},
	"3H": {
		Face:  "three",
		Suit:  "hearts",
		Value: 3,
	},
	"2H": {
		Face:  "two",
		Suit:  "hearts",
		Value: 2,
	},
	"AH": {
		Face:  "ace",
		Suit:  "hearts",
		Value: 1,
	},
}
