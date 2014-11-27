package main

import (
	"fmt"
)

const (
	eulerMax = 100
)

func sumOfSquares(sumChan chan int64, max int64) {
	var i int64
	for i = 1; i <= max; i++ {
		sumChan <- i * i
	}
	close(sumChan)
}

func squareOfSums(squareChan chan int64, max int64) {
	var i int64
	sum := int64(0)
	for i = 1; i <= max; i++ {
		sum += i
	}
	squareChan <- sum * sum
	close(squareChan)
}

func squareMinusSum(max int64) (squareMinusSum int64) {
	sumChan := make(chan int64)
	squareChan := make(chan int64)

	go squareOfSums(squareChan, max)
	go sumOfSquares(sumChan, max)

	var sumDone, squareDone bool
	for {
		if sumDone && squareDone {
			return
		}
		select {
		case receive, ok := <-sumChan:
			if !ok {
				sumDone = true
			} else {
				squareMinusSum -= receive
			}
		case receive, ok := <-squareChan:
			if !ok {
				squareDone = true
			} else {
				squareMinusSum += receive
			}
		}
	}
}

func main() {
	answer := squareMinusSum(eulerMax)
	fmt.Printf("The square of the sums minus the sums of the squares of 1 to %d is %d\n", eulerMax, answer)
}
