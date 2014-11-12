package main

import (
    "fmt"
)

const (
    limit = 4e6
)

func emitFib(fibChan chan int) {

    // Problem says start with [1, 2]. I prefer the more classical [0, 1]
    current := 1
    prior := 0

    // Go doesn't have a "while" - only "for"
    for current <= limit {

        // Emit current
        fibChan <- current

        // Step
        new := current + prior
        current, prior = new, current
    }
    close(fibChan)
}

func processFib(fibChan chan int) (sum int) {
    for {
        val, ok := <-fibChan
        if !ok {
            // closed channel
            return
        }

        if val % 2 == 0 {
            sum += val
        }
    }
}

func calculate() int {
    fibChan := make(chan int)
    go emitFib(fibChan)

    return processFib(fibChan)

}

func main() {
    sum := calculate()
    fmt.Printf("Sum of even fibonnaci numbers is %d.\n", sum)
}
