package collatz
var memory map[int][]int

func CollatzWithMemory(start int) []int {
	if memory == nil {
		memory = make(map[int][]int)
	}

	if chain, exists := memory[start]; exists {
		return chain
	}

	var store []int
	store = append(store, int(start))
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
		} else {
			start = (3 * start) + 1
		}
		store = append(store, int(start))
	}

	memory[store[0]] = store

	return store
}
