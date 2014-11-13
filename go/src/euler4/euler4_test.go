package main

import (
    "testing"
)

func TestIntegralPalindromesArePalindromes(t *testing.T) {
    palindromes := [...]int{11, 101, 111, 1221, 12321, 1234321}
    for _, n := range palindromes {
        if !isPalindrome(n) {
            t.Errorf("Number %d incorrectly classified as non-palindrome", n)
        }
    }
}

func TestNonpalindromesAreNotPalindromes(t *testing.T) {
    palindromes := [...]int{12, 201, 102, 112, 17221, 3141, 12345678}
    for _, n := range palindromes {
        if isPalindrome(n) {
            t.Errorf("Number %d incorrectly classified as palindrome", n)
        }
    }
}

