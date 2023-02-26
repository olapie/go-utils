package utils

import (
	"context"
	"testing"
)

type ID int64

func TestGetLogin(t *testing.T) {
	t.Run("int64ToID", func(t *testing.T) {
		ctx := WithLogin(context.TODO(), int64(1))
		id := GetLogin[ID](ctx)
		if id != ID(1) {
			t.FailNow()
		}
	})

	t.Run("int64ToString", func(t *testing.T) {
		ctx := WithLogin(context.TODO(), int64(1))
		id := GetLogin[string](ctx)
		if id != "1" {
			t.FailNow()
		}
	})
}
