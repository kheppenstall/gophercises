package deck

import "testing"

func TestSuitConstants(t *testing.T) {
	suits := map[Suit]int{
		Spade:   0,
		Diamond: 1,
		Club:    2,
		Heart:   3,
		Joker:   4,
	}

	for suit, exp := range suits {
		if int(suit) != exp {
			t.Errorf("Expected %v, got %v", exp, suit)
		}
	}
}

func TestSuitString(t *testing.T) {
	suits := map[Suit]string{
		Spade:   "Spades",
		Diamond: "Diamonds",
		Club:    "Clubs",
		Heart:   "Hearts",
		Joker:   "Joker",
	}

	for suit, exp := range suits {
		s := suit.String()
		if s != exp {
			t.Errorf("Expected %v, got %v", exp, s)
		}
	}

}
