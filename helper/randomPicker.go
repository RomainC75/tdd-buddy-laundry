package helper

import "math/rand/v2"

func GetRandomInArray[T any](arr []T) T {
	ln := len(arr)
	return arr[rand.IntN(ln)]
}
