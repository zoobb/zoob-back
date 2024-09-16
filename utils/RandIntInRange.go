package utils

import "math/rand"

func RandIntInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}
