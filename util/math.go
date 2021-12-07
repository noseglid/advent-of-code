package util

import "math"

func Factorial(i int) int {
	result := 1
	for i > 0 {
		result *= i
		i--
	}

	return result
}

// PermInt calls f with each permutation of a.
func PermInt(a []int, f func([]int)) {
	permInt(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func permInt(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permInt(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permInt(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func Absolute(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MaxIntList(a []int) int {
	m := math.MinInt64
	for _, n := range a {
		if n > m {
			m = n
		}
	}
	return m
}

func MinIntList(a []int) int {
	m := math.MaxInt64
	for _, n := range a {
		if n < m {
			m = n
		}
	}
	return m
}
