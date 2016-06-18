package main

import (
	"fmt"
	"math"
)

func calculatePythagTripletSum(sum int) (int, error) {
	/*
	   See blog post for derivation - but basically, calculate required `b` given
	   an `a`, then see if a is an integer and if it matches the parameters.
	*/

	// float64 so we can use math standard library
	var a, b float64
	n := float64(sum)

	for b = 1; b < n/2; b++ {
		// Calculate required `a`
		a = (2*b*n - n*n) / (2*b - 2*n)

		// is `a` an integer?
		// (built-in modulus doesn't work on floats)
		if math.Mod(a, 1) == 0 {
			return int(a * b * (n - a - b)), nil
		}
	}
	return 0, fmt.Errorf("Sum %f does not appear to have a triplet", n)
}

func main() {
	eulerSum := 1000
	answer, err := calculatePythagTripletSum(eulerSum)
	if err != nil {
		panic("No answer found")
	}
	fmt.Printf("The product of the digits of a pythagorean triplet of %d is %d \n", eulerSum, answer)
}
