package utils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("IntArray", func(t *testing.T) {
		a := []int{1, 2, 3, -9, 10, 1, 101}
		ReverseSlice(a)
		MustEqualT(t, []int{101, 1, 10, -9, 3, 2, 1}, a)

		a = []int{1}
		ReverseSlice(a)
		MustEqualT(t, []int{1}, a)

		a = []int{}
		ReverseSlice(a)
		MustEqualT(t, []int{}, a)

		a = []int{1, 3}
		ReverseSlice(a)
		MustEqualT(t, []int{3, 1}, a)
	})

	t.Run("StringArray", func(t *testing.T) {
		a := []string{"a", "b", "c", "d"}
		ReverseSlice(a)
		MustEqualT(t, []string{"d", "c", "b", "a"}, a)
	})
}
