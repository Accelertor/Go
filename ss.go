package main

func main() {
	s := 0
	for i := 2; i < 1000000; i++ {
		if isprime(i) {
			s += i
		}
	}
	println(s)
}
