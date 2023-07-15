package utils

import (
	"context"
	"fmt"
	"log"
	"reflect"
)

type keyType int

// Context keys
const (
	keyStart keyType = iota
	keyLogin
	keySudo
	keyLogger
	keyRequestInfo

	keyEnd
)

// Deprecated: use context.WithoutCancel() instead
func DetachContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	for k := keyStart; k < keyEnd; k++ {
		if v := ctx.Value(k); v != nil {
			newCtx = context.WithValue(newCtx, k, v)
		}
	}
	return newCtx
}

func GetLogin[T comparable](ctx context.Context) T {
	v := ctx.Value(keyLogin)
	if v == nil {
		var zero T
		return zero
	}

	if actual, ok := v.(T); ok {
		return actual
	}

	actualVal := reflect.ValueOf(v)
	var expect T

	switch actualVal.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		if reflect.ValueOf(expect).Kind() == reflect.String {
			reflect.ValueOf(&expect).Elem().SetString(fmt.Sprint(v))
			return expect
		}
	default:
		break
	}

	expectType := reflect.TypeOf(expect)
	if actualVal.Type().ConvertibleTo(expectType) {
		defer func() {
			if msg := recover(); msg != nil {
				log.Printf("[sugar/v2/contexts] GetLogin: %v\n", msg)
			}
		}()
		reflect.ValueOf(&expect).Elem().Set(reflect.ValueOf(v).Convert(expectType))
	}
	return expect
}

func WithLogin[T comparable](ctx context.Context, v T) context.Context {
	var zero T
	if v == zero {
		if ctx.Value(keyLogin) == nil {
			return ctx
		}
		return context.WithValue(ctx, keyLogin, nil)
	}
	return context.WithValue(ctx, keyLogin, v)
}

func WithSudo(ctx context.Context) context.Context {
	return context.WithValue(ctx, keySudo, true)
}

func IsSudo(ctx context.Context) bool {
	b, _ := ctx.Value(keySudo).(bool)
	return b
}

type requestContextInfo struct {
	AppID         string
	Authorization string
	ClientID      string
	ServiceID     string
	TraceID       string
	TestFlag      bool
}

type RequestContextBuilder interface {
	Build() context.Context
	WithAppID(v string) RequestContextBuilder
	WithAuthorization(v string) RequestContextBuilder
	WithClientID(v string) RequestContextBuilder
	WithServiceID(v string) RequestContextBuilder
	WithTraceID(v string) RequestContextBuilder
	WithTestFlag(v bool) RequestContextBuilder
}

type requestContextBuilderImpl struct {
	ctx  context.Context
	info requestContextInfo
}

func NewRequestContextBuilder(ctx context.Context) RequestContextBuilder {
	return &requestContextBuilderImpl{ctx: ctx}
}

func (b *requestContextBuilderImpl) Build() context.Context {
	return context.WithValue(b.ctx, keyRequestInfo, &b.info)
}

func (b *requestContextBuilderImpl) WithAppID(v string) RequestContextBuilder {
	b.info.AppID = v
	return b
}

func (b *requestContextBuilderImpl) WithAuthorization(v string) RequestContextBuilder {
	b.info.Authorization = v
	return b
}

func (b *requestContextBuilderImpl) WithClientID(v string) RequestContextBuilder {
	b.info.ClientID = v
	return b
}

func (b *requestContextBuilderImpl) WithServiceID(v string) RequestContextBuilder {
	b.info.ServiceID = v
	return b
}

func (b *requestContextBuilderImpl) WithTraceID(v string) RequestContextBuilder {
	b.info.TraceID = v
	return b
}

func (b *requestContextBuilderImpl) WithTestFlag(v bool) RequestContextBuilder {
	b.info.TestFlag = v
	return b
}

func GetAppID(ctx context.Context) string {
	info, _ := ctx.Value(keyRequestInfo).(*requestContextInfo)
	if info == nil {
		return ""
	}
	return info.AppID
}

func GetAuthorization(ctx context.Context) string {
	info, _ := ctx.Value(keyRequestInfo).(*requestContextInfo)
	if info == nil {
		return ""
	}
	return info.Authorization
}

func GetTraceID(ctx context.Context) string {
	info, _ := ctx.Value(keyRequestInfo).(*requestContextInfo)
	if info == nil {
		return ""
	}
	return info.TraceID
}

func GetServiceID(ctx context.Context) string {
	info, _ := ctx.Value(keyRequestInfo).(*requestContextInfo)
	if info == nil {
		return ""
	}
	return info.ServiceID
}

func GetClientID(ctx context.Context) string {
	info, _ := ctx.Value(keyRequestInfo).(*requestContextInfo)
	if info == nil {
		return ""
	}
	return info.ClientID
}

func IsTest(ctx context.Context) bool {
	info, _ := ctx.Value(keyRequestInfo).(*requestContextInfo)
	if info == nil {
		return false
	}
	return info.TestFlag
}
