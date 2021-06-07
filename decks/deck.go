package decks

import (
	"errors"
	"math/rand"
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/jordanjordanb1/cards-api/cache"
)

type Deck struct {
	Id           string `json:"deck_id"`
	Shuffled     bool   `json:"shuffled"`
	Remaining    int    `json:"remaining"`
	UndrawnCards []Card `json:"cards,omitempty"`
	DrawnCards   []Card `json:"drawn_cards,omitempty"`
}

type Card struct {
	Index int    `json:"-"`
	Face  string `json:"face"`
	Suit  string `json:"suit"`
	Value int    `json:"value"`
}

// Creates a new deck of cards with no custom selection
func NewDeck(shuffle bool) Deck {
	var _cards []Card
	uuid := uuid.New().String()
	shuffled := false

	// Gets all cards without map key
	_cards = mapToArray(BasicDeck)

	if shuffle {
		// Shuffle cards
		_cards = ShuffleDeck(_cards)
		shuffled = true
	} else {
		// Make sure cards are in order
		sortDeck(_cards)
	}

	deck := Deck{Id: uuid, Shuffled: shuffled, Remaining: 52, UndrawnCards: _cards}

	return deck
}

// Creates a new deck based on selection passed as an argument
func NewCustomDeck(shuffle bool, cardSelection []string) Deck {
	var _cards []Card
	uuid := uuid.New().String()
	shuffled := false

	// Loops through array of passed card types
	for _, cardHash := range cardSelection {
		foundCard, found := BasicDeck[cardHash]

		// Checks if card type is valid
		if found {
			_cards = append(_cards, foundCard)
		}
	}

	if shuffle {
		// Shuffle cards
		_cards = ShuffleDeck(_cards)
		shuffled = true
	} else {
		// Make sure cards are in order
		sortDeck(_cards)
	}

	deck := Deck{Id: uuid, Shuffled: shuffled, Remaining: len(_cards), UndrawnCards: _cards}

	return deck
}

// Saves deck in the in-memory cache
func (deck *Deck) Save() *Deck {
	// Deletes old item if set by any chance
	cache.Cache.Del(deck.Id)

	cache.Cache.Set(deck.Id, deck, 1)

	return deck
}

// Get deck from cache
func GetDeckById(id string) (*Deck, error) {
	deck, found := cache.Cache.Get(id)

	// If deck not found, return error
	if !found {
		return nil, errors.New("No deck found with ID of " + id)
	}

	return deck.(*Deck), nil
}

// Shuffles deck
func ShuffleDeck(cards []Card) []Card {
	// Required to seed randomizer
	rand.Seed(time.Now().UnixNano())

	// Shuffles cards
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards
}

// Sorts deck of cards based on it's index
func sortDeck(cards []Card) {
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].Index < cards[j].Index
	})
}

// Flattens a card map to an array
func mapToArray(mapToFlatten map[string]Card) []Card {
	var acc []Card

	for _, value := range mapToFlatten {
		acc = append(acc, value)
	}

	return acc
}
