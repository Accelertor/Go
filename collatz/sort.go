package collatz

import (
	"fmt"
	"sort"
)

type InputOutPut struct {
	digit  int
	length int
}

func Sort() {
	var pair []InputOutPut
	for i := 1; i < 1500000; i++ {
		output := ArrayLength(CollatzWithMemory(i)) //storying length
		pair = append(pair, InputOutPut{digit: i, length: output})
	}
	sort.Slice(pair, func(i, j int) bool {
		return pair[i].length < pair[j].length
	})

	fmt.Println("Sorted Digit-Length")
	for _, pairs := range pair {
		fmt.Printf("Digit:%d,Length of Array:%d\n", pairs.digit, pairs.length)
	}

}
