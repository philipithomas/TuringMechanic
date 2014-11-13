package main

import (
    "fmt"
    "strconv"
)

func isPalindrome(n int) bool {
    // Convert integer to byte array (via string)
    str := []byte(strconv.Itoa(n))
    for i:=0; i < (len(str)/2); i++ {
        if str[i] != str[(len(str)-i-1) % len(str)] {
            return false
        }
    }
    return true
}

func main() {
}
