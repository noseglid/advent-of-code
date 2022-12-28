package util

func Contains[T comparable](list []T, el T) bool {
	for _, l := range list {
		if l == el {
			return true
		}
	}

	return false
}

func RemoveByValue[T comparable](list []T, el T) ([]T, bool) {
	for i, l := range list {
		if el == l {
			list[i], list[len(list)-1] = list[len(list)-1], list[i]
			return list[0 : len(list)-1], true
		}
	}
	return list, false
}
