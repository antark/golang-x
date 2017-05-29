package random_util

import (
	"fmt"
	"math"
	"testing"
)

func TestRandomHit(t *testing.T) {
	odd := 15
	limit := 1000000
	hitCount := 0
	for i := 0; i < limit; i++ {
		if Hit(odd) {
			hitCount++
		}
	}
	fmt.Printf("%v%% odds of %v times, count is: %v\n", odd, limit, hitCount)
	fmt.Printf("error is %v%%\n", math.Abs(float64(limit*odd/100)-float64(hitCount))*float64(100)/float64(limit*odd/100))
}
