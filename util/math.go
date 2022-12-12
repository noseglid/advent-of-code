package util

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Factorial(i int) int {
	result := 1
	for i > 0 {
		result *= i
		i--
	}

	return result
}

func Perm(a []interface{}, f func([]interface{})) {
	perm(a, f, 0)
}

func perm(a []interface{}, f func([]interface{}), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func PermString(a []string, f func([]string)) {
	permString(a, f, 0)
}

func permString(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permString(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permString(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
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

func Absolute[T constraints.Integer | constraints.Float](i T) T {
	if i < 0 {
		return -i
	}
	return i
}

func Min[T constraints.Integer | constraints.Float](i T, j ...T) T {
	min := i
	for _, t := range j {
		if t < min {
			min = t
		}
	}
	return min
}

func Max[T constraints.Integer | constraints.Float](i T, j ...T) T {
	max := i
	for _, t := range j {
		if t > max {
			max = t
		}
	}
	return max
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
