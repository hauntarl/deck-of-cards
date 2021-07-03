// Deck module provides the basic structures and constants which can be utilized
// to create a deck of cards for a given card game, eg. Blackjack.
package deck

import "sort"

// New is a constructor for deck of cards.
func New(options ...func([]Card) []Card) []Card {
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
