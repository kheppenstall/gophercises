package deck

// Rank enum represents rank of card
type Rank uint8

// Two, Three... represent the 13 ranks
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (r Rank) String() string {
	m := rankMap()
	return m[r]
}

func rankMap() map[Rank]string {
	return map[Rank]string{
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
}
