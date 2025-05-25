package utils

import "slices"

func NextId[T any](data []T, id func(T) int) int {
	if len(data) == 0 {
		return 1
	}
	last := slices.MaxFunc(data, func(a, b T) int {
		return id(a) - id(b)
	})
	return id(last) + 1
}
