
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
    lcmOut = 1
    for i:= 1; i<=max; i++ {
        lcmOut  = lcm(lcmOut, i)
    }
    return
}

func main() {
    max := calculate(eulerMax)
    fmt.Printf("The max palindrome formed by the multplication of two three-digit numbers is %d\n", max)
}
