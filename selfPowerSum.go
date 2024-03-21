package main

import (
	"fmt"
	"math/big"
)

func selfpow(n int) {
	 // Define your limit 'n' here
	sum := new(big.Int)

	for i := 1; i <= n; i++ {
		power := new(big.Int).Exp(big.NewInt(int64(i)), big.NewInt(int64(i)), nil)
		sum.Add(sum, power)
	}

	// Convert the result to a string
	resultStr := sum.String()

	// Get the last 10 digits
	last10Digits := resultStr[len(resultStr)-10:]

	fmt.Printf("Sum of numbers raised to their own power (1^1 + 2^2 + ... + %d^%d) (last 10 digits): %s\n", n, n, last10Digits)
}
