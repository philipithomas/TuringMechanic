package main

import (
    "fmt"
)

var (
    eulerLim int = 2000000
)


func primeSum(lim int)  (sum int64) {
    // False means prime at output
    sieve := make([]bool, lim)

    // Not a prime
    sieve[0] = true
    sieve[1] = true

    prime := 2
    for {
        sum += int64(prime)
        // Get rid of multiples of prime
        for i:=prime; i<lim; i=i+prime {
            if i % prime == 0 {
                sieve[i] = true
            }
        }

        // Find next prime
        for {
            prime++
            if prime == lim {
                return
            }
            if sieve[prime] == false {
                // found next prime
                break
            }
        }
    }
}

func main() {
    answer := primeSum(eulerLim)
    fmt.Printf("The sum of prime number below %d is %d\n", eulerLim, answer)
}
