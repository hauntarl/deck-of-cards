//go:generate stringer -type=Suit,Rank -output card_string.go
// Deck module provides the basic structures and constants which can be utilized
// to create a deck of cards for a given card game, eg. Blackjack.
package deck

import "fmt"

// New is a constructor for deck of cards.
func New() []Card {
	cards := make([]Card, 0, len(suits)*int(maxRank))
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}

// Card is a uniquely identifiable combination of Suit and Rank in a deck.
type Card struct {
	Suit
	Rank
}

// String returns a human readable representation of a card.
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%v of %vs", c.Rank, c.Suit)
}

// Suit represents type of the card in a deck.
type Suit uint8

// The suits have been intentionally put in a certain order i.e. the order that
// cards are typically sorted in a brand new deck.
//  Joker // is a special case.
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

// suits allow us to iterate over all the suits excluding Joker.
var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank represents value of the card in a deck.
type Rank uint8

// All the ranks have values that correspond to the actually what they are.
//  Ace // might have special point values, they are the exception.
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// minRank and maxRank allow us to iterate over all the ranks while creating a
// deck of cards.
const (
	minRank = Ace
	maxRank = King
)
