package board

import "testing"

const A1_OUTPUT string = "" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"X-------"

const A2_OUTPUT string = "" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"X-------" + "\n" +
	"--------"

const A8_OUTPUT string = "" +
	"X-------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------"

const H8_OUTPUT string = "" +
	"-------X" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------"

const H1_OUTPUT string = "" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"-------X"

const E4_OUTPUT string = "" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"----X---" + "\n" +
	"--------" + "\n" +
	"--------" + "\n" +
	"--------"

func TestSquareStringA1(t *testing.T) {
	square := GetSquareFromString("A1")
	if square.String() != A1_OUTPUT {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", A1_OUTPUT, square.String())
	}
}

func TestSquareStringA2(t *testing.T) {
	square := GetSquareFromString("A2")
	if square.String() != A2_OUTPUT {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", A2_OUTPUT, square.String())
	}
}

func TestSquareStringA8(t *testing.T) {
	square := GetSquareFromString("A8")
	if square.String() != A8_OUTPUT {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", A8_OUTPUT, square.String())
	}
}

func TestSquareStringH1(t *testing.T) {
	square := GetSquareFromString("H1")
	if square.String() != H1_OUTPUT {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", H1_OUTPUT, square.String())
	}
}

func TestSquareStringH8(t *testing.T) {
	square := GetSquareFromString("H8")
	if square.String() != H8_OUTPUT {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", H8_OUTPUT, square.String())
	}
}

func TestSquareStringE4(t *testing.T) {
	square := GetSquareFromString("E4")
	if square.String() != E4_OUTPUT {
		t.Fatalf("\nExpected: \n%s\nActual: \n%s", E4_OUTPUT, square.String())
	}
}

func TestSquareRank(t *testing.T) {
	square := GetSquareFromString("A1")
	if square.GetRank() != 1 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 1, square.GetRank())
	}

	square = GetSquareFromString("E4")
	if square.GetRank() != 4 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 4, square.GetRank())
	}

	square = GetSquareFromString("H8")
	if square.GetRank() != 8 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 8, square.GetRank())
	}

	square = GetSquareFromString("C8")
	if square.GetRank() != 8 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 8, square.GetRank())
	}

	square = GetSquareFromString("A2")
	if square.GetRank() != 2 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 2, square.GetRank())
	}
}

func TestSquareFile(t *testing.T) {
	square := GetSquareFromString("A1")
	if square.GetFile() != 1 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 1, square.GetFile())
	}

	square = GetSquareFromString("E4")
	if square.GetFile() != 5 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 5, square.GetFile())
	}

	square = GetSquareFromString("H8")
	if square.GetFile() != 8 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 8, square.GetFile())
	}

	square = GetSquareFromString("C8")
	if square.GetFile() != 3 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 3, square.GetFile())
	}

	square = GetSquareFromString("A2")
	if square.GetFile() != 1 {
		t.Fatalf("\nExpected: \n%d\nActual: \n%d", 1, square.GetFile())
	}
}
