package main

import "fmt"

func printplis(x string) {
	lenh := len(x)
	for i := 0; i < lenh; i++ {
		for j := 0; j < lenh; j++ {
			if i == lenh/2 {
				fmt.Printf("%c", x[j])
			} else if j == lenh/2 {
				fmt.Printf("%c", x[i])
			} else {
				fmt.Print(" ")
			}
		}
		println("")
	}
}
