package main

import (
	"fmt"
)

const (
	eulerMax = 20
)

// Greatest common divisor
func gcd(i, j int) int {
	// Euclid's algorithm
	k := i % j
	if k == 0 {
		return j
	}
	return gcd(j, k)
}

// Least common multiple
func lcm(i, j int) int {
	// from LCM wikipedia page
	return i * j / gcd(i, j)
}

func calculate(max int) (lcmOut int) {
	// What is the LCM of all integers between 1 and max, inclusive?
	lcmOut = 1
	for i := 1; i <= max; i++ {
		lcmOut = lcm(lcmOut, i)
	}
	return
}

func main() {
	lcmOut := calculate(eulerMax)
	fmt.Printf("The LCM of numbers from 1 to %d is %d\n", eulerMax, lcmOut)
}
