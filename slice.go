package utils

import (
	"reflect"
)

func UniqueSlice[E comparable](a []E) []E {
	m := make(map[E]struct{}, len(a))
	l := make([]E, 0, len(a))
	for _, v := range a {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		l = append(l, v)
	}
	return l
}

func ReverseArray(a any) bool {
	a = Indirect(a)
	if a == nil {
		return false
	}
	v := reflect.ValueOf(a)
	if v.IsNil() || !v.IsValid() || v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return false
	}

	for i, j := 0, v.Len()-1; i < j; i, j = i+1, j-1 {
		vi, vj := v.Index(i), v.Index(j)
		tmp := vi.Interface()
		if !vi.CanSet() {
			return false
		}
		vi.Set(vj)
		vj.Set(reflect.ValueOf(tmp))
	}
	return true
}

func RemoveElement[E comparable](a []E, v E) []E {
	for i, e := range a {
		if e == v {
			a = append(a[:i], a[i+1:]...)
			break
		}
	}
	return a
}

func SliceToSet[E comparable](a []E) map[E]bool {
	m := make(map[E]bool)
	for _, v := range a {
		m[v] = true
	}
	return m
}

func MustTransformSlice[E1 any, E2 any](a []E1, f func(e E1) E2) []E2 {
	res := make([]E2, len(a))
	for i, e := range a {
		res[i] = f(e)
	}
	return res
}

func TransformSlice[E1 any, E2 any](a []E1, f func(e E1) (E2, error)) ([]E2, error) {
	res := make([]E2, len(a))
	for i, e := range a {
		e2, err := f(e)
		if err != nil {
			return nil, err
		}
		res[i] = e2
	}
	return res, nil
}

func CastToIntSlice[T ~int | ~int32 | ~int16 | ~int8](a []T) []int {
	res := make([]int, len(a))
	for i, v := range a {
		res[i] = int(v)
	}
	return res
}

func CastFromIntSlice[T ~int](a []int) []T {
	res := make([]T, len(a))
	for i, v := range a {
		res[i] = T(v)
	}
	return res
}

func CastToInt16Slice[T ~int16 | ~int8 | ~int](a []T) []int16 {
	res := make([]int16, len(a))
	for i, v := range a {
		res[i] = int16(v)
	}
	return res
}

func CastFromInt16Slice[T ~int16](a []int16) []T {
	res := make([]T, len(a))
	for i, v := range a {
		res[i] = T(v)
	}
	return res
}

func CastToInt32Slice[T ~int32 | ~int16 | ~int8 | ~int](a []T) []int32 {
	res := make([]int32, len(a))
	for i, v := range a {
		res[i] = int32(v)
	}
	return res
}

func CastFromInt32Slice[T ~int32](a []int32) []T {
	res := make([]T, len(a))
	for i, v := range a {
		res[i] = T(v)
	}
	return res
}

func CastToInt64Slice[T ~int64 | ~int16 | ~int32 | ~int | ~int8](a []T) []int64 {
	res := make([]int64, len(a))
	for i, v := range a {
		res[i] = int64(v)
	}
	return res
}

func CastFromInt64Slice[T ~int64](a []int64) []T {
	res := make([]T, len(a))
	for i, v := range a {
		res[i] = T(v)
	}
	return res
}

func CastToStringSlice[T ~string](a []T) []string {
	res := make([]string, len(a))
	for i, v := range a {
		res[i] = string(v)
	}
	return res
}

func CastFromStringSlice[T ~string](a []string) []T {
	res := make([]T, len(a))
	for i, v := range a {
		res[i] = T(v)
	}
	return res
}
