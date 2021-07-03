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
type Option func(cards []Card) []Card

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

// Sort creates a user-defined sorting option, whose returned value can be
// passed to the deck constructor and have cards sorted in the specified way.
func Sort(less func(cards []Card) func(i, j int) bool) Option {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

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

// Jokers option allows you to add jokers in your deck of cards.
func Jokers(n int) Option {
	return func(cards []Card) []Card {
		for ; n > 0; n-- {
			// each joker has been assigned a different rank for you to
			// differentiate between them
			cards = append(cards, Card{Suit: Joker, Rank: Rank(n)})
		}
		return cards
	}
}

// Filter option will remove set of cards which match the given predicate.
func Filter(predicate func(Card) bool) Option {
	return func(cards []Card) []Card {
		end := len(cards) - 1
		for i := 0; i <= end; {
			if predicate(cards[i]) {
				cards[end], cards[i] = cards[i], cards[end]
				end--
			} else {
				i++
			}
		}
		return cards[: end+1 : end+1]
	}
}
