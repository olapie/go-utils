package utils

import (
	"testing"
)

func TestMax(t *testing.T) {
	t.Run("N0", func(t *testing.T) {
		v := Max[int]()
		MustEqualT(t, 0, v)
	})

	t.Run("N1", func(t *testing.T) {
		v := Max(10)
		MustEqualT(t, 10, v)
	})

	t.Run("N2", func(t *testing.T) {
		v := Max(-0.3, 10.9)
		MustEqualT(t, 10.9, v)
	})

	t.Run("N3", func(t *testing.T) {
		v := Max(-0.3, 10.9, 3.8)
		MustEqualT(t, 10.9, v)
	})
}
