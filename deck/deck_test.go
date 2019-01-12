package deck

import "testing"

func TestNew(t *testing.T) {
	deck := New()

	exp := 54
	if len(deck) != exp {
		t.Errorf("Expected %v, got %v", exp, len(deck))
	}
}
