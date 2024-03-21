package main

import "fmt"

func indexPrime(lastIndex int) {
	primecount := 0
	currentNumber := 2
	for {
		if isprime(currentNumber) {
			primecount++

			if primecount == lastIndex {
				break
			}

		}

		currentNumber++
	}

	fmt.Printf("the prime numbver on %d is %d \n", lastIndex, currentNumber)
}
