package main

import (
    "fmt"
)

var (
    eulerPrime int = 10001
)

func primeNum(n int)  int {
    // Start with known 2 as prime
    var i int
    i = 2
    primes := []int{i,}
    for {
        i++
        isPrime := true
        for _, prime := range primes {
            if i % prime == 0 {
                isPrime = false
                break
            }
        }

        if isPrime {
            primes = append(primes, i)
            if len(primes) == n {
                return i
            }
        }
    }
}

func main() {
    answer := primeNum(eulerPrime)
    fmt.Printf("Prime number %d is %d\n", eulerPrime, answer)
}


