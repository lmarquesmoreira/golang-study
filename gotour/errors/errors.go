package main

import (
	"fmt"
	"math"
)

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.2f", e)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN(), errNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	z := 9999.
	fmt.Println(sqrt(z * -1))
	fmt.Println(sqrt(z))
}
