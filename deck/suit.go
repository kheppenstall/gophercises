package deck

// Suit enum represents suit of card
type Suit uint8

// Spade, Diamond, Club, Heart, Joker represent the five suits
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

func (s Suit) String() string {
	m := suitMap()
	return m[s]
}

func suitMap() map[Suit]string {
	return map[Suit]string{
		Spade:   "Spades",
		Diamond: "Diamonds",
		Club:    "Clubs",
		Heart:   "Hearts",
		Joker:   "Joker",
	}
}
