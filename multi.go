package main

import (
	"fmt"
	"math"
)

func multi(x, y float64) float64 {
	r := math.Log10(x) + math.Log10(y)
	u := math.Pow(10, r)
	return u
}
func myfacto(limit float64) float64 {
	pro := 1.0
	for i := 1; i <= int(limit); i++ {
		pro = (multi(float64(pro), float64(i)))

	}
	return pro

}

func su(x float64) {
	s := 0
	F2I := math.Float64bits(x)
	fmt.Printf("%d", F2I)
	for F2I != 0 {
		s += int(F2I) % 10
		F2I /= 10
	}
	//fmt.Printf("%d", s)
}
