package utils

import (
	"time"
)

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64 | time.Duration
}

func Abs[T Number](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

func Sum[T Number](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}
