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

func RemoveByIndex[T any](list []T, index int) []T {
	return append(list[:index], list[index+1:]...)
}

func Unique[T comparable](list []T) []T {
	set := map[T]struct{}{}
	for _, el := range list {
		set[el] = struct{}{}
	}

	unq := []T{}
	for el := range set {
		unq = append(unq, el)
	}

	return unq
}

func Repeat[T any](el T, count int) []T {
	r := []T{}
	for i := 0; i < count; i++ {
		r = append(r, el)
	}
	return r
}
