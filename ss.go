package main

import (
	"fmt"
	"math"
)

// Function to generate pandigital permutations of length n
func generatePandigitalPermutations(n int) [][]int {
	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = n - i
	}
	return generatePermutations(digits)
}

// Function to generate permutations of a set of integers
func generatePermutations(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}} // Base case: Return an empty permutation
	}

	var permutations [][]int

	for i, num := range nums {
		// Create a copy of nums without the current element
		remaining := make([]int, len(nums)-1)
		copy(remaining[:i], nums[:i])
		copy(remaining[i:], nums[i+1:])

		// Recursively generate permutations of the remaining elements
		subPermutations := generatePermutations(remaining)

		// Append the current element to the beginning of each permutation
		for _, perm := range subPermutations {
			permutations = append(permutations, append([]int{num}, perm...))
		}
	}

	return permutations
}

func convertToInt(x []int) int {
	a := 0
	for _, v := range x {
		a = a*10 + v
	}
	return a
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	// Define the number of digits for pandigital numbers
	n := 7 // Change this to find largest n-digit pandigital prime

	// Generate pandigital permutations of length n
	permutations := generatePandigitalPermutations(n)

	// Find the largest pandigital prime
	largestPrime := -1
	for _, perm := range permutations {
		number := convertToInt(perm)
		if isPrime(number) {
			largestPrime = int(math.Max(float64(largestPrime), float64(number)))
		}
	}

	// Print the largest pandigital prime found
	if largestPrime != -1 {
		fmt.Println("Largest", n, "-digit pandigital prime:", largestPrime)
	} else {
		fmt.Println("No", n, "digit pandigital prime found")
	}
}
