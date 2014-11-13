package main

import (
	"fmt"
	"math"
)

const (
	x = 600851475143
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		// include i!=n for formality to pass test for 2
		if n%i == 0 && i != n {
			return false
		}
	}
	return true
}

func calculate(x int) int {
	// Start with largest possible prime number
	// "int" will take the floor of the returned float
	n := int(math.Sqrt(float64(x)))
	for {
		// is factor?
		if x%n == 0 {
			// is prime?
			if isPrime(n) {
				// Max prime factor!
				return n
			}
		}
		n--
	}
}

func main() {
	max := calculate(x)
	fmt.Printf("The maximum number prime factor of %d is %d\n", x, max)
}
