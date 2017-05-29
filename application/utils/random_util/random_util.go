package random_util

import (
	"math/rand"
	"time"
)

func Hit(odd int) bool {
	if odd < 0 || odd > 100 {
		panic("invalid odd parameter. odd ∈ [0, 100]")
	}

	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(100)
	if value < odd {
		return true
	}
	return false
}
