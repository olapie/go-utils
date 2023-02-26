package utils

import (
	"fmt"
)

func WrapError(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}
	a = append(a, err)
	return fmt.Errorf(format+":%w", a...)
}

// CauseError returns the root cause error
func CauseError(err error) error {
	for {
		u, ok := err.(interface{ Unwrap() error })
		if !ok {
			break
		}
		err = u.Unwrap()
	}
	return err
}

func GetErrorCode(err error) int {
	if s, ok := err.(interface{ Code() int }); ok {
		return s.Code()
	}

	if s, ok := err.(interface{ GetCode() int }); ok {
		return s.GetCode()
	}

	// ------------
	// int32

	if s, ok := err.(interface{ Code() int32 }); ok {
		return int(s.Code())
	}

	if s, ok := err.(interface{ GetCode() int32 }); ok {
		return int(s.GetCode())
	}

	if s, ok := err.(interface{ StatusCode() int }); ok {
		return s.StatusCode()
	}

	if s, ok := err.(interface{ GetStatusCode() int }); ok {
		return s.GetStatusCode()
	}

	if s, ok := err.(interface{ Status() int }); ok {
		return s.Status()
	}

	if s, ok := err.(interface{ GetStatus() int }); ok {
		return s.GetStatus()
	}

	// ------------
	// int32

	if s, ok := err.(interface{ StatusCode() int32 }); ok {
		return int(s.StatusCode())
	}

	if s, ok := err.(interface{ GetStatusCode() int32 }); ok {
		return int(s.GetStatusCode())
	}

	if s, ok := err.(interface{ Status() int32 }); ok {
		return int(s.Status())
	}

	if s, ok := err.(interface{ GetStatus() int32 }); ok {
		return int(s.GetStatus())
	}

	return 0
}
