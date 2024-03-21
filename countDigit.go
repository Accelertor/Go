package main

func countDigit(num float64) int {
	if num == 0 {
		return 0
	}
	return 1 + countDigit(num/10)
}
