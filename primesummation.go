package main

func sp() {
	count := 2000000
	sum := 0
	for i := 0; i < count-1; i++ {
		if isprime(i) {
			sum += i
		}

	}
	print(sum)
}
