package main

import (
    "testing"
)

func TestPythagTripletReturnsTrue(t *testing.T) {
    if isPythagTriplet(3, 4, 5) != true {
        t.Error("Failed to detect a pythagorean triplet")
    }
}

func TestNonPythagTripletReturnsFalse(t *testing.T) {
    if isPythagTriplet(1, 1, 1) == true {
        t.Error("False positive for pythag triplet")
    }
    if isPythagTriplet(3, 5, 4) == true {
        t.Error("False positive for pythag triplet")
    }
}

func TestIncrementingRulesEquilateral(t *testing.T) {
    a, b, c := 2, 2, 2  // input
    x, y, z := 2, 2, 3 // Expected output
    alpha, beta, gamma, err := incrementTriplet(a, b, c)
    if alpha != x || beta != y || gamma != z || err != nil {
        t.Errorf("Unexpected increment for %d, %d, %d - %d, %d, %d", a, b, c, alpha, beta, gamma)
    }
}

func TestIncrementingRulesPythag(t *testing.T) {
    a, b, c := 3, 4, 5 // input
    x, y, z := 4, 5, 6 // Expected output
    alpha, beta, gamma, err := incrementTriplet(a, b, c)
    if alpha != x || beta != y || gamma != z || err != nil {
        t.Errorf("Unexpected increment for %d, %d, %d - %d, %d, %d", a, b, c, alpha, beta, gamma)
    }
}

func TestNotTriangleLargeC(t *testing.T) {
    a, b, c := 2, 2, 4 // input
    x, y, z := 2, 3, 4 // Expected output
    alpha, beta, gamma, err := incrementTriplet(a, b, c)
    if alpha != x || beta != y || gamma != z || err != nil {
        t.Errorf("Unexpected increment for %d, %d, %d - %d, %d, %d", a, b, c, alpha, beta, gamma)
    }
}

// Small theta means the angle alpha means c^2 > a^2 + b^2
func TestIncrementingRulesSmallTheta(t *testing.T) {
    a, b, c := 3, 4, 6 // input
    x, y, z := 4, 4, 6 // Expected output
    alpha, beta, gamma, err := incrementTriplet(a, b, c)
    if alpha != x || beta != y || gamma != z || err != nil {
        t.Errorf("Unexpected increment for %d, %d, %d - %d, %d, %d", a, b, c, alpha, beta, gamma)
    }
}

func TestIncrementingRulesLargeTheta(t *testing.T) {
    a, b, c := 4, 4, 5 // input
    x, y, z := 4, 4, 6 // Expected output
    alpha, beta, gamma, err := incrementTriplet(a, b, c)
    if alpha != x || beta != y || gamma != z || err != nil {
        t.Errorf("Unexpected increment for %d, %d, %d - %d, %d, %d", a, b, c, alpha, beta, gamma)
    }
}

func TestEulerExample(t *testing.T) {
    sum := 3 + 4 + 5
    out, err := calculatePythagTripletSum(sum)
    expected := 3*4*5
    if err != nil {
        t.Error("Unexpected error")
    }
    if out != expected {
        t.Errorf("Incorrect answer")
    }
}

func TestTripletExample(t *testing.T) {
    sum := 5 + 12 + 13
    out, err := calculatePythagTripletSum(sum)
    expected := 5*12*13
    if err != nil {
        t.Error("Unexpected error")
    }
    if out != expected {
        t.Errorf("Incorrect answer")
    }
}
