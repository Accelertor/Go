package main

import (
	"fmt"
	"math/big"
)

func ps(expo int) {
	// Create a big integer representing the base (2)
	base := big.NewInt(2)

	// Create a big integer representing the exponent
	exponent := big.NewInt(int64(expo))

	// Calculate the result of 2^expo
	result := new(big.Int).Exp(base, exponent, nil)

	// Initialize a big integer to store the sum of digits
	sum := new(big.Int)

	// Create a copy of the result to work with
	temp := new(big.Int).Set(result)

	// Create big integers for 10 and 0
	ten := big.NewInt(10)
	zero := big.NewInt(0)

	// Loop to extract and sum the digits
	for temp.Cmp(zero) > 0 {
		// Divide temp by 10 and get the quotient and remainder
		quotient, digit := temp.DivMod(temp, ten, new(big.Int))

		// Add the remainder (digit) to the sum
		sum.Add(sum, digit)

		// Set temp to the quotient for the next iteration
		temp.Set(quotient)
	}

	// Print the result of 2^expo
	fmt.Println("2^", expo, ":", result.String())

	// Print the sum of digits
	fmt.Println("Sum of digits:", sum.String())
}
