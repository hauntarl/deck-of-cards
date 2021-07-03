// Deck module provides the basic structures and constants which can be utilized
// to create a deck of cards for a given card game, eg. Blackjack.
package deck

import (
	"math/rand"
	"sort"
)

// New is a constructor for deck of cards.
func New(options ...Option) []Card {
	cards := make([]Card, 0, len(suits)*int(maxRank))
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, fn := range options {
		cards = fn(cards)
	}
	return cards
}

// Option represents the signature of function which can be passed as options.
type Option func([]Card) []Card

// Shuffle is a functional option for deck constructor which will return a
// shuffled deck of cards.
//
// PS: there is no point passing Shuffle and Sort options together in any order.
func Shuffle(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

// Sort creates a user-defined sorting option, whose returned value can be
// passed to the deck constructor and have cards sorted in the specified way.
func Sort(less func(cards []Card) func(i, j int) bool) Option {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// DefaultSort is a functional option that can be passed to the constructor for
// sorting the deck of cards in a default way.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Less function provides a default way of comparing two different cards
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return value(cards[i]) < value(cards[j])
	}
}

// value calculates the absolute value of any given card which is later
// utilized when sorting the deck using a default implementation.
func value(c Card) int { return int(c.Suit)*int(maxRank) + int(c.Rank) }
