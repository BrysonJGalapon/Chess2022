package board

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}

func Pow2(x int) int {
	return x * x
}

func Xor(b1, b2 bool) bool {
	return (b1 && !b2) || (!b1 && b2)
}
