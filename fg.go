package main

import "fmt"

const mod = 10000

func ex48() {
	var n, k int
	var p2, p1, p0, s2, s1, s0, t2, t1, t0 uint

	// s = 0
	s2, s1, s0 = 0, 0, 0
	for n = 1; n <= 1000; n++ {
		// p = 1
		p2, p1, p0 = 0, 0, 1
		for k = 1; k <= n; k <<= 1 {
		}
		for k >>= 1; k > 0; k >>= 1 {
			// p *= p
			t0 = p0 * p0
			t1 = 2*p0*p1 + t0/mod
			t2 = 2*p0*p2 + p1*p1 + t1/mod
			p0 = t0 % mod
			p1 = t1 % mod
			p2 = t2 % mod
			if n&k != 0 {
				// p *= n
				t0 = p0 * uint(n)
				t1 = p1*uint(n) + t0/mod
				t2 = p2*uint(n) + t1/mod
				p0 = t0 % mod
				p1 = t1 % mod
				p2 = t2 % mod
			}
		}
		// s += p
		t0 = p0 + s0
		t1 = p1 + s1 + t0/mod
		t2 = p2 + s2 + t1/mod
		s0 = t0 % mod
		s1 = t1 % mod
		s2 = t2 % mod
	}

	fmt.Printf("ex48: %02d%04d%04d\n", s2%100, s1, s0)
}

func main() {
	ex48()
}
