package player

import "math/rand"

func GetRandomInt(arr []int) int {
	i := rand.Intn(len(arr))
	return arr[i]
}

func Shuffle(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		j := i + rand.Intn(n-i)

		// swap i and j
		tmp := arr[j]
		arr[j] = arr[i]
		arr[i] = tmp
	}
}
