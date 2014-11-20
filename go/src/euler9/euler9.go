package main

import (
    "fmt"
)

func isPythagTriplet(a, b, c int) bool {
    return a*a + b*b == c*c
}

func calculatePythagTripletSum(sum int) (int, error) {
    a, b, c := 1, 1, 1 // Starting point
    var err error
    for {
        fmt.Printf("%d, %d, %d\n", a, b, c)
        if a + b + c < sum {
            a, b, c, err = incrementTriplet(a, b, c)
            if err != nil {
                // panic because this shouldn't happen
                panic(err)
            }
            continue
        }
        if isPythagTriplet(a, b, c) {
            return a*b*c, nil
        }
        if a + b + c > sum {
            // sum isn't possible
            return 0, fmt.Errorf("No pythagorean triplet sums to %s", sum)
        }
        a, b, c, err = incrementTriplet(a, b, c)
    }
}

func incrementTriplet(alpha, beta, gamma int) (a, b, c int, err error) {
    // input and output variables must be different :-(
    a, b, c = alpha, beta, gamma

    // case triplet:
    if isPythagTriplet(a, b, c) {
        a++; b++; c++
        return
    }

    // case not triangle small C
    if a + b <= c {
        for a + b <= c {
            if a == b {
                b++
            } else {
                a++
            }
        }
        return
    }

    // case equilateral
    if a == b && b == c {
        c++
        return
    }

    // case small theta
    if a*a + b*b < c*c {
        if b > a {
            a++
        } else {
            b++
        }
        return
    }

    // case small theta
    if a*a + b*b > c*c {
        c++
        return
    }


    return 0, 0, 0, fmt.Errorf("Uncaught case %d, %d, %d", alpha, beta, gamma)
}

func main() {
    eulerSum := 100
    answer, err := calculatePythagTripletSum(eulerSum)
    if err != nil {
        fmt.Printf("%s", err)
    }
    fmt.Printf("The  product of the digits of a pythagorean triplet of %d is %d \n", eulerSum, answer)
}
