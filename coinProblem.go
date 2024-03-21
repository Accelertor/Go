package main

import "fmt"

func findCoins(x int) {
	arr := []int{10, 5, 2, 1}
	num := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		for x >= arr[i] {
			num[arr[i]]++
			x -= arr[i]
		}
	}

	for key, value := range num {
		fmt.Printf("Coin Denomination: %d, Count: %d\n", key, value)
	}
}
