package collatz

import "fmt"

func Collats(start int) {
	var store []int
	fmt.Printf("sequence for %d : ", start)
	store = append(store, int(start))
	//THIS IS WHERE WE ARE CALCULATING COLLATZ
	for start != 1 {

		if start%2 == 0 {
			start = start / 2
			/*println(" ", start)
			println(" | |")*/
			store = append(store, int(start))
		} else {
			start = (3 * start) + 1
			/*println(" ", start)
			println("  | |")*/
			store = append(store, int(start))
		}

	}
	//THIS IS WHERE I'M PRINTING IT
	for i := 0; i <= len(store)-1; i++ {
		if store[i] == 1 {
			println("1")
			break
		}
		print(store[i], "->")
	}
	println("  ")
}
