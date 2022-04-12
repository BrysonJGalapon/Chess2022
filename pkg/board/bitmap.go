package board

type BitMap uint64

var bitmapToSquare map[BitMap]Square = make(map[BitMap]Square)

func init() {
	var square Square = 1
	var bitmap BitMap = 1
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			bitmapToSquare[bitmap] = square
			square <<= 1
			bitmap <<= 1
		}
	}
}

func (b BitMap) ToSquare() Square {
	return bitmapToSquare[b]
}
