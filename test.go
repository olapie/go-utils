package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func MustEqualT(t testing.TB, expected, result any, args ...any) {
	diff := cmp.Diff(expected, result)
	if diff == "" {
		return
	}
	msg := fmt.Sprint(args...)
	if msg != "" {
		t.Fatal(msg, "\n", diff)
	} else {
		t.Fatal(diff)
	}
}

func ShouldEqualT(t testing.TB, expected, result any, args ...any) {
	diff := cmp.Diff(expected, result)
	if diff == "" {
		return
	}
	msg := fmt.Sprint(args...)
	if msg != "" {
		t.Error(msg, "\n", diff)
	} else {
		t.Error(diff)
	}
}

func MustDiffT(t testing.TB, expected, result any, args ...any) {
	diff := cmp.Diff(expected, result)
	if diff == "" {
		t.Fatal(args...)
	}
}

func ShouldDiffT(t testing.TB, expected, result any, args ...any) {
	diff := cmp.Diff(expected, result)
	if diff == "" {
		t.Error(args...)
	}
}

func MustTrueT(t testing.TB, b bool, args ...any) {
	if !b {
		args = append([]any{"Expected true, got false"}, args...)
		t.Fatal(args...)
	}
}

func ShouldTrueT(t testing.TB, b bool, args ...any) {
	if !b {
		args = append([]any{"Expected true, got false"}, args...)
		t.Fatal(args...)
	}
}

func MustFalseT(t testing.TB, b bool, args ...any) {
	if b {
		args = append([]any{"Expected false, got true"}, args...)
		t.Fatal(args...)
	}
}

func ShouldFalseT(t testing.TB, b bool, args ...any) {
	if b {
		args = append([]any{"Expected false, got true"}, args...)
		t.Fatal(args...)
	}
}

func MustErrorT(t testing.TB, err error, args ...any) {
	if err != nil {
		args = append([]any{err}, args...)
		t.Fatal(args...)
	}
}

func ShouldErrorT(t testing.TB, err error, args ...any) {
	if err != nil {
		args = append([]any{err}, args...)
		t.Error(args...)
	}
}

func MustNotErrorT(t testing.TB, err error, args ...any) {
	if err == nil {
		t.Fatal(args...)
	}
}

func ShouldNotErrorT(t testing.TB, err error, args ...any) {
	if err == nil {
		t.Error(args...)
	}
}

func isEmpty(i any) bool {
	if i == nil {
		return true
	}

	v := reflect.ValueOf(i)
	k := v.Kind()
	switch {
	case k == reflect.Pointer:
		return v.IsNil()
	case k == reflect.Map || k == reflect.Slice:
		return v.IsNil() && v.Len() == 0
	case k == reflect.String:
		return v.String() == ""
	case k == reflect.Bool:
		return !v.Bool()
	case v.CanInt():
		return v.Int() == 0
	case v.CanUint():
		return v.Uint() == 0
	case v.CanFloat():
		return v.Float() == 0
	default:
		return false
	}
}

func MustEmptyT(t testing.TB, i any, args ...any) {
	if isEmpty(i) {
		return
	}
	s := fmt.Sprint(i)
	if len(s) > 512 {
		s = s[:512] + "..."
	}
	t.Fatalf("must be empty: %s %s", s, fmt.Sprint(args...))
}

func ShouldEmptyT(t testing.TB, i any, args ...any) {
	if isEmpty(i) {
		return
	}
	s := fmt.Sprint(i)
	if len(s) > 512 {
		s = s[:512] + "..."
	}
	t.Errorf("should be empty: %s %s", s, fmt.Sprint(args...))
}

func MustNotEmptyT(t testing.TB, i any, args ...any) {
	if isEmpty(i) {
		s := fmt.Sprint(i)
		if len(s) > 512 {
			s = s[:512] + "..."
		}
		t.Fatalf("must not be empty: %s %s", s, fmt.Sprint(args...))
	}
}

func ShouldNotEmptyT(t testing.TB, i any, args ...any) {
	if isEmpty(i) {
		s := fmt.Sprint(i)
		if len(s) > 512 {
			s = s[:512] + "..."
		}
		t.Errorf("should not be empty: %s %s", s, fmt.Sprint(args...))
	}
}
