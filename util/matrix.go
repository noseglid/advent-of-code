package util

func Flip2D[T any](m [][]T) [][]T {
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}

	return m
}

func Rotate2D[T any](m [][]T) [][]T {
	m = Flip2D(m)
	for i := 0; i < len(m); i++ {
		for j := 0; j < i; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	return m
}
