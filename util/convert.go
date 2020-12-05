package util

import (
	"fmt"
	"strconv"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("'%s' can not be interpreted as an int", s))
	}
	return i
}
