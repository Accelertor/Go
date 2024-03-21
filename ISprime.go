package main

/*THIS CODE OPATATE ON PRINIPLE THAT MOST PRIME ABOVE 3 ARE SOME MULTIPLE OF 6 -1*/
func isprime(num int) bool {
	if num <= 1 {
		return false
	}

	if num == 2 || num == 3 {
		return true
	}

	if num%2 == 0 || num%3 == 0 {
		return false
	}

	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}

	return true
}
