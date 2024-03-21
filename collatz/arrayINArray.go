package collatz

func ArrayLength(inputA []int) int {
	var distin []int
	lesn := 0
	for _, value := range inputA {
		distin = append(distin, value)
	}
	lesn = len(distin)
	return lesn
}
