package main

import (
	"fmt"
	"math"
)

func triplet() {
	a := 0.0
	b := 0.0
	c := 0.0
	for a = 1.0; a <= 1000.0; a++ {
		for b = a; b <= 1000.0; b++ {
			c = 1000.0 - a - b
			if math.Pow(a, 2)+math.Pow(b, 2) == math.Pow(c, 2) {
				fmt.Printf("a=%f,n=%f,c=%f\n", a, b, c)
				println(a * b * c)
			}
		}
	}
}
