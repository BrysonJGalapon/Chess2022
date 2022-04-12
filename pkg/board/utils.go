package board

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Pow2(x int) int {
	return x * x
}

func Xor(b1, b2 bool) bool {
	return (b1 && !b2) || (!b1 && b2)
}

func NumSetBits(n BitMap) int {
	count := 0
	for n != 0 {
		if n&1 != 0 {
			count += 1
		}
		n >>= 1
	}
	return count
}

func NumSteps(srcSquare, dstSquare Square) int {
	startRow := srcSquare.GetRow()
	startCol := srcSquare.GetCol()
	endRow := dstSquare.GetRow()
	endCol := dstSquare.GetCol()

	return Max(Abs(endRow-startRow), Abs(endCol-startCol))
}
