package util

func Factorial(i int) int {
	result := 1
	for i > 0 {
		result *= i
		i--
	}

	return result
}
