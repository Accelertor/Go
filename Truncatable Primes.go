package main

import (
	"math"
)

func rightTruncatablePrime(num int) bool {
	for num != 0 {
		if isprime(num) == false {
			return false
		}
		num /= 10
	}

	return true
}

func leftTruncatablePrime(num int) bool {
	D := countDigit(float64(num))
	i := 0
	for num != 0 {
		if !isprime(int(num)) {
			return false
		}
		num %= int(math.Pow10(D - i))
		i++
	}
	return true
}
func main() {
	i := 0
	s := 0
	for j := 9; i < 11; j++ {
		if rightTruncatablePrime(j) && leftTruncatablePrime(j) {
			s += j
			i++
		}

	}

	println(s)

}
