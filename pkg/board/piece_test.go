package board

import "testing"

func TestIsValidMovement(t *testing.T) {
	var p *Piece = BLACK_QUEEN
	err := p.IsValidMovement(GetSquareFromString("D8"), GetSquareFromString("G3"))
	if err == nil {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", "non-nil-error", err)
	}
}
