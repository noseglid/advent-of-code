package util

func RangeInt(n int) []int {
	var l []int
	for i := 0; i < n; i++ {
		l = append(l, i)
	}

	return l
}
