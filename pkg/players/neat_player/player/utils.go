package player

import "math/rand"

func getRandomInt(arr []int) int {
	i := rand.Intn(len(arr))
	return arr[i]
}
