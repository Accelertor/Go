package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(n int) int {
	largest := -1

	// Divide by 2 until n becomes odd
	for n%2 == 0 {
		largest = 2
		n /= 2
	}

	// Now n is odd, so we can start checking odd factors
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}

	// If n is still greater than 2, then it's a prime number
	if n > 2 {
		largest = n
	}

	return largest
}

func mainzz() {
	num := 600851475143 // Replace this with the number you want to find the largest prime factor of
	result := largestPrimeFactor(num)
	fmt.Printf("The largest prime factor of %d is %d\n", num, result)
}
