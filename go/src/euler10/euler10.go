package main

import (
    "fmt"
)

var (
    eulerLim int = 2*10e6
    chanLim int = 10
)

func emit(ch chan<- int) {
    for i:=2; ; i++ {
        ch <- i
    }
}

func filter(in <-chan int, out chan<- int, prime int) {
    for {
        i := <-in
        if i%prime != 0 {
            out <- i
        }
    }
}


func primeSum(lim int)  (sum int64) {
    chEmit := make(chan int)
    go emit(chEmit)
    for {
        prime := <-chEmit
        fmt.Printf("%d\n", prime)
        if prime > lim {
            return
        }
        sum += int64(prime)
        chReceive := make(chan int)
        go filter(chEmit, chReceive, prime)
        chEmit = chReceive
    }
}

func main() {
    answer := primeSum(eulerLim)
    fmt.Printf("The sum of prime number below %d is %d\n", eulerLim, answer)
}
