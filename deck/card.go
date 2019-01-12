package deck

import "fmt"

// Card has a suit and rank
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%v of %v", c.Rank.String(), c.Suit.String())
}

func newCard(s Suit, r Rank) Card {
	return Card{
		Suit: s,
		Rank: r,
	}
}

func newJoker() Card {
	return newCard(Joker, 0)
}
