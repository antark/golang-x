package math_ext

import (
	"fmt"
	"testing"
)

func TestMathExt(t *testing.T) {
	a, b := 10, 20
	fmt.Printf("max(%v,%v) == %v\n", a, b, MaxInt(a, b))
	fmt.Printf("min(%v,%v) == %v\n", a, b, MinInt(a, b))

	a, b = 20, 10
	fmt.Printf("max(%v,%v) == %v\n", a, b, MaxInt(a, b))
	fmt.Printf("min(%v,%v) == %v\n", a, b, MinInt(a, b))

	a, b = 10, 10
	fmt.Printf("max(%v,%v) == %v\n", a, b, MaxInt(a, b))
	fmt.Printf("min(%v,%v) == %v\n", a, b, MinInt(a, b))
}
