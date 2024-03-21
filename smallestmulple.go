package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func maaaain() {
	result := 1
	for i := 2; i <= 10; i++ {
		result = lcm(result, i)
	}
	fmt.Println(result)
}
