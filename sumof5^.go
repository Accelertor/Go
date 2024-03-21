package main

import (
	"math"
	"math/big"
	"strconv"
)

func sumOfDigits(n *big.Int) int {
	sum := 0
	str := n.String()
	for _, digit := range str {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
	}
	return sum}
func extractDigits(x int) []int {
	a := []int{}
	for x != 0 {
		a = append(a, x%10)
		x /= 10
	}
	return a
}

func is5thPower(x int) bool {
	a := extractDigits(x)
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += int(math.Pow(float64(a[i]), 5))
	}
	if sum == x {
		return true
	} else {
		return false
	}
}
