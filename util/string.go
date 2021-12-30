package util

func Reverse(s string) string {
	var rev []rune
	for i := len(s) - 1; i >= 0; i-- {
		rev = append(rev, rune(s[i]))
	}
	return string(rev)
}
