package board

import "testing"

const STANDARD_BOARD_STRING string = "" +
	"rnbqkbnr" + "\n" +
	"pppppppp" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"PPPPPPPP" + "\n" +
	"RNBQKBNR"

func TestBoardStringStandard(t *testing.T) {
	var b Board = Standard()
	if b.String() != STANDARD_BOARD_STRING {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", STANDARD_BOARD_STRING, b.String())
	}
}
