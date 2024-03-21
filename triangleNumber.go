package main

/*A TRIANGLE NUMBER IS MADE BY ADDING ALL INT FROM 1 TO THAT NUMBER */
func TriNumber(limit int) int {

	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i
	}
	return sum
}

func PrimeDivisorFequency(n int) map[int]int {
	fMap := make(map[int]int)
	for i := 2; i <= n; i++ {
		if isprime(i) && n%i == 0 {
			fMap[i]++
			n /= i
			i--
		}
	}
	return fMap
}

func FerqToCount(x int) int {
	count := 1
	f := PrimeDivisorFequency(x)
	for _, freq := range f {
		count *= (freq + 1)
	}
	return count
}
func check() {
	for i := 1; ; i++ {
		g := FerqToCount(TriNumber(i))
		if g > 500 {
			println(i)
			println(TriNumber(i))
			break
		}

	}
}
