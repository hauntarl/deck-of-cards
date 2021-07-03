// Deck module provides the basic structures and constants which can be utilized
// to create a deck of cards for a given card game, eg. Blackjack.
package deck

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
