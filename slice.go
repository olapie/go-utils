package utils

import (
	"reflect"
)

func CloneSlice[T any](a []T) []T {
	res := make([]T, len(a))
	copy(res, a)
	return res
}

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

func ReverseSlice[E comparable](a []E) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
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

func IndexOfSlice[E comparable](a []E, v E) int {
	for i, e := range a {
		if e == v {
			return i
		}
	}
	return -1
}

func FilterSlice[E any](a []E, filter func(e E) bool) []E {
	res := make([]E, 0, len(a)/2)
	for _, v := range a {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}
