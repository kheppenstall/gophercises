package deck

import "testing"

func TestNewCard(t *testing.T) {
	c := newCard(Heart, Jack)

	if c.Suit != Heart {
		t.Errorf("Expected %v, got %v", Heart, c.Suit)
	}

	if c.Rank != Jack {
		t.Errorf("Expected %v, got %v", Jack, c.Rank)
	}
}

func TestNewJoker(t *testing.T) {
	j := newJoker()

	if j.Suit != Joker {
		t.Errorf("Expected %v, got %v", Joker, j.Suit)
	}
}
func TestCardString(t *testing.T) {
	card := Card{Spade, Two}
	s := card.String()
	exp := "Two of Spades"

	if s != exp {
		t.Errorf("Expected %v, got %v", exp, s)
	}
}
func TestCardStringAsJoker(t *testing.T) {
	card := Card{Joker, 0}
	s := card.String()
	exp := "Joker"

	if s != exp {
		t.Errorf("Expected %v, got %v", exp, s)
	}
}
