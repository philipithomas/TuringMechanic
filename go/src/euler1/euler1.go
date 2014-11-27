package main

import (
	"fmt"
)

const (
	limit = 1000
)

func calculate() (sum int) {
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return
}

func main() {
	sum := calculate()
	fmt.Printf("The sum of multiples of 3 or 5 below %d is %d.\n", limit, sum)
}
