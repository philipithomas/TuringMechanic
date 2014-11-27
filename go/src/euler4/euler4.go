package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	lower = 100
	upper = 999
)

func isPalindrome(n int) bool {
	// Convert integer to byte array (via string)
	str := []byte(strconv.Itoa(n))
	for i := 0; i < (len(str) / 2); i++ {
		if str[i] != str[(len(str)-i-1)%len(str)] {
			return false
		}
	}
	return true
}

func calculate() (max int, err error) {
	for i := lower; i <= upper; i++ {
		for j := lower; j <= upper; j++ {
			test := i * j
			if isPalindrome(test) {
				if max < test {
					max = test
				}
			}
		}
	}
	if max == 0 {
		err = errors.New("No value found")
	}
	return
}

func main() {
	max, err := calculate()
	if err != nil {
		panic(err)
	}
	fmt.Printf("The max palindrome formed by the multplication of two three-digit numbers is %d\n", max)
}
