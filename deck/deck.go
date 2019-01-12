package deck

// New instantiates a deck
func New() []Card {
	deck := []Card{}
	for r := 1; r <= 13; r++ {
		rank := Rank(r)
		for s := 0; s <= 3; s++ {
			suit := Suit(s)
			card := newCard(suit, rank)
			deck = append(deck, card)
		}
	}

	for j := 0; j < 2; j++ {
		joker := newJoker()
		deck = append(deck, joker)
	}

	return deck
}
