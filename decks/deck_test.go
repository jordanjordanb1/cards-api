package decks

import (
	"sort"
	"testing"

	"github.com/google/uuid"
)

// Tests creation of a new deck with shuffle set to false
func TestNewDeckNoShuffle(t *testing.T) {
	deck := NewDeck(false)

	// Checks for remaining cards
	if deck.Remaining != 52 {
		t.Errorf("NewDeck(shuffle = false) returned an incorrect remaining amount of cards. Expected: '52' and Recieved: '%d'", deck.Remaining)
	}

	// Checks for correct shuffled value
	if deck.Shuffled != false {
		t.Errorf("NewDeck(shuffle = false) returned an incorrect shuffled value. Expected: 'false' and Recieved: '%t'", deck.Shuffled)
	}

	// Checks to see if UUID is valid
	_, err := uuid.Parse(deck.Id)

	if err != nil {
		t.Errorf("NewDeck(shuffle = false) returned an invalid UUID. Recieved: '%s' Error Message: '%s'", deck.Id, err)
	}

	// Checks to see if the drawn cards array is empty
	if len(deck.DrawnCards) != 0 {
		t.Errorf("NewDeck(shuffle = false) returned an invalid length of drawn cards. Recieved: '%d' Expected: '0'", len(deck.DrawnCards))
	}

	// Checks to see if the undrawn cards array has the same length as the BasicDeck map
	if len(deck.UndrawnCards) != len(BasicDeck) {
		t.Errorf("NewDeck(shuffle = false) returned an invalid length for the undrawn cards field. Recieved: '%d' Expected: '52'", len(deck.UndrawnCards))
	}

	// Checks to see if deck is sorted
	isSorted := isSorted(deck.UndrawnCards)

	if !isSorted {
		t.Errorf("NewDeck(shuffle = false) was checked to be sorted when it should've been as 'shuffle' is set to false")
	}
}

// Tests creation of a new deck with shuffle set to true
func TestNewDeckWithShuffle(t *testing.T) {
	deck := NewDeck(true)

	// Checks for remaining cards
	if deck.Remaining != 52 {
		t.Errorf("NewDeck(shuffle = true) returned an incorrect remaining amount of cards. Expected: '52' and Recieved: '%d'", deck.Remaining)
	}

	// Checks for correct shuffled value
	if deck.Shuffled != true {
		t.Errorf("NewDeck(shuffle = true) returned an incorrect shuffled value. Expected: 'true' and Recieved: '%t'", deck.Shuffled)
	}

	// Checks to see if UUID is valid
	_, err := uuid.Parse(deck.Id)

	if err != nil {
		t.Errorf("NewDeck(shuffle = true) returned an invalid UUID. Recieved: '%s' Error Message: '%s'", deck.Id, err)
	}

	// Checks to see if the drawn cards array is empty
	if len(deck.DrawnCards) != 0 {
		t.Errorf("NewDeck(shuffle = true) returned an invalid length of drawn cards. Recieved: '%d' Expected: '0'", len(deck.DrawnCards))
	}

	// Checks to see if the undrawn cards array has the same length as the BasicDeck map
	if len(deck.UndrawnCards) != len(BasicDeck) {
		t.Errorf("NewDeck(shuffle = true) returned an invalid length for the undrawn cards field. Recieved: '%d' Expected: '52'", len(deck.UndrawnCards))
	}

	// Checks to see if deck is sorted
	isSorted := isSorted(deck.UndrawnCards)

	if isSorted {
		t.Errorf("NewDeck(shuffle = true) was checked to be sorted when it should'nt have been as 'shuffle' is set to true")
	}
}

// Tests creation of a new custom deck with shuffle set to false
func TestNewCustomDeckNoShuffle(t *testing.T) {
	selection := []string{"7S", "8S", "KS", "JD", "AC"}
	deck := NewCustomDeck(false, selection)

	// Checks for remaining cards
	if deck.Remaining != len(selection) {
		t.Errorf("NewCustomDeck(shuffle = false) returned an incorrect remaining amount of cards. Expected: '%d' and Recieved: '%d'", len(selection), deck.Remaining)
	}

	// Checks for correct shuffled value
	if deck.Shuffled != false {
		t.Errorf("NewCustomDeck(shuffle = false) returned an incorrect shuffled value. Expected: 'false' and Recieved: '%t'", deck.Shuffled)
	}

	// Checks to see if UUID is valid
	_, err := uuid.Parse(deck.Id)

	if err != nil {
		t.Errorf("NewCustomDeck(shuffle = false) returned an invalid UUID. Recieved: '%s' Error Message: '%s'", deck.Id, err)
	}

	// Checks to see if the drawn cards array is empty
	if len(deck.DrawnCards) != 0 {
		t.Errorf("NewCustomDeck(shuffle = false) returned an invalid length of drawn cards. Recieved: '%d' Expected: '0'", len(deck.DrawnCards))
	}

	// Checks to see if deck is sorted
	isSorted := isSorted(deck.UndrawnCards)

	if !isSorted {
		t.Errorf("NewCustomDeck(shuffle = false) was checked to be sorted when it should've been as 'shuffle' is set to false")
	}
}

// Tests creation of a new custom deck with shuffle set to false
func TestNewCustomDeckWithShuffle(t *testing.T) {
	selection := []string{"9S", "3S", "QS", "2D", "2C", "7D", "AS"}
	deck := NewCustomDeck(true, selection)

	// Checks for remaining cards
	if deck.Remaining != len(selection) {
		t.Errorf("NewCustomDeck(shuffle = true) returned an incorrect remaining amount of cards. Expected: '%d' and Recieved: '%d'", len(selection), deck.Remaining)
	}

	// Checks for correct shuffled value
	if deck.Shuffled != true {
		t.Errorf("NewCustomDeck(shuffle = true) returned an incorrect shuffled value. Expected: 'true' and Recieved: '%t'", deck.Shuffled)
	}

	// Checks to see if UUID is valid
	_, err := uuid.Parse(deck.Id)

	if err != nil {
		t.Errorf("NewCustomDeck(shuffle = true) returned an invalid UUID. Recieved: '%s' Error Message: '%s'", deck.Id, err)
	}

	// Checks to see if the drawn cards array is empty
	if len(deck.DrawnCards) != 0 {
		t.Errorf("NewCustomDeck(shuffle = true) returned an invalid length of drawn cards. Recieved: '%d' Expected: '0'", len(deck.DrawnCards))
	}

	// Checks to see if deck is sorted
	isSorted := isSorted(deck.UndrawnCards)

	if isSorted {
		t.Errorf("NewCustomDeck(shuffle = true) was checked to be sorted when it shouldn't have been as 'shuffle' is set to true")
	}
}

func isSorted(cards []Card) bool {
	return sort.SliceIsSorted(cards[:], func(i, j int) bool {
		return cards[i].Index < cards[j].Index
	})
}
