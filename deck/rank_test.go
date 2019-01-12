package deck

import "testing"

func TestRankConstants(t *testing.T) {
	ranks := map[Rank]int{
		Ace:   1,
		Two:   2,
		Three: 3,
		Four:  4,
		Five:  5,
		Six:   6,
		Seven: 7,
		Eight: 8,
		Nine:  9,
		Ten:   10,
		Jack:  11,
		Queen: 12,
		King:  13,
	}

	for rank, exp := range ranks {
		if int(rank) != exp {
			t.Errorf("Expected %v, got %v", exp, rank)
		}
	}
}

func TestRankString(t *testing.T) {
	ranks := map[Rank]string{
		Ace:   "Ace",
		Two:   "Two",
		Three: "Three",
		Four:  "Four",
		Five:  "Five",
		Six:   "Six",
		Seven: "Seven",
		Eight: "Eight",
		Nine:  "Nine",
		Ten:   "Ten",
		Jack:  "Jack",
		Queen: "Queen",
		King:  "King",
	}

	for r, exp := range ranks {
		s := r.String()
		if s != exp {
			t.Errorf("Expected %v, got %v", exp, s)
		}
	}
}
