package board

import "fmt"

type Square uint64

var SQUARE_STRINGS [64]string = [64]string{
	"A8", "B8", "C8", "D8", "E8", "F8", "G8", "H8",
	"A7", "B7", "C7", "D7", "E7", "F7", "G7", "H7",
	"A6", "B6", "C6", "D6", "E6", "F6", "G6", "H6",
	"A5", "B5", "C5", "D5", "E5", "F5", "G5", "H5",
	"A4", "B4", "C4", "D4", "E4", "F4", "G4", "H4",
	"A3", "B3", "C3", "D3", "E3", "F3", "G3", "H3",
	"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2",
	"A1", "B1", "C1", "D1", "E1", "F1", "G1", "H1",
}

var coordToSquare [8][8]Square = [8][8]Square{{}, {}, {}, {}, {}, {}, {}, {}}
var stringToSquare map[string]Square = make(map[string]Square)
var squareToBitMap map[Square]BitMap = make(map[Square]BitMap)

func init() {
	var square Square
	var bitmap BitMap

	square = 1
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			coordToSquare[r][c] = square
			square <<= 1
		}
	}

	square = 1
	for _, s := range SQUARE_STRINGS {
		stringToSquare[s] = square
		square <<= 1
	}

	square = 1
	bitmap = 1
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			squareToBitMap[square] = bitmap
			square <<= 1
			bitmap <<= 1
		}
	}
}

func (s *Square) GetX() int {
	// TODO
	panic("Unimplemented")
}

func (s *Square) GetY() int {
	// TODO
	panic("Unimplemented")
}

func (s *Square) ToBitMap() BitMap {
	ret := squareToBitMap[*s]
	return ret
}

func (s *Square) String() string {
	ret := ""
	var square Square = 1
	for i := 0; i < 64; i++ {
		if *s == square {
			ret += "X"
		} else {
			ret += "-"
		}

		if (i%8 == 7) && (i != 63) {
			ret += "\n"
		}

		square <<= 1
	}

	return ret
}

func (s *Square) GetName() string {
	var square Square = 1
	for _, squareString := range SQUARE_STRINGS {
		if *s == square {
			return squareString
		}
		square <<= 1
	}

	panic(fmt.Sprintf("could not find name for square: %d", s))
}

func GetSquareFromString(s string) Square {
	return stringToSquare[s]
}

func GetSquareFromCoord(row, col int) Square {
	return coordToSquare[row][col]
}
