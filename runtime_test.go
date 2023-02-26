package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestIndirectKind(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		k := IndirectKind(nil)
		MustEqualT(t, reflect.Invalid, k)
	})

	t.Run("Struct", func(t *testing.T) {
		var p time.Time
		k := IndirectKind(p)
		MustEqualT(t, reflect.Struct, k)
	})

	t.Run("PointerToStruct", func(t *testing.T) {
		var p *time.Time
		k := IndirectKind(p)
		MustEqualT(t, reflect.Struct, k)
	})

	t.Run("PointerToPointerToStruct", func(t *testing.T) {
		var p **time.Time
		k := IndirectKind(p)
		MustEqualT(t, reflect.Struct, k)
	})

	t.Run("Map", func(t *testing.T) {
		var p map[string]any
		k := IndirectKind(p)
		MustEqualT(t, reflect.Map, k)
	})

	t.Run("PointerToMap", func(t *testing.T) {
		var p map[string]any
		k := IndirectKind(p)
		MustEqualT(t, reflect.Map, k)
	})
}
