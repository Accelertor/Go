package main

func findProperDivisor(x int) int {
	sum := 0
	for i := 1; i <= x/2; i++ {
		if x%i == 0 {
			sum += i
		}
	}
	return sum
}

func isAmicable(x, y int) bool {
	if findProperDivisor(x) == y && findProperDivisor(y) == x {
		return true
	}
	return false
}

func findAmicable() []int {
	a := []int{}
	for i := 1; i <= 10000; i++ {
		sum := findProperDivisor(i)
		if sum != i && sum <= 10000 && findProperDivisor(sum) == i {
			a = append(a, i)
		}
	}
	return a
}
