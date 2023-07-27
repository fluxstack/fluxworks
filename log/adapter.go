package log

import (
	"context"
	"github.com/fluxstack/fluxworks/types"
)

type Adapter interface {
	Log(level Level, fields types.M) error
}

type adapter struct {
	fields  types.M
	adapter Adapter
	prefix  []interface{}
	ctx     context.Context
}

func (w *adapter) Log(level Level, fields types.M) error {
	kvs := make(types.M, len(w.fields)+len(fields))
	for k, v := range w.fields {
		kvs[k] = v
	}
	for k, v := range fields {
		kvs[k] = v
	}
	return w.adapter.Log(level, kvs)
}

func merge(fields ...types.M) types.M {
	kvs := types.M{}
	for _, f := range fields {
		for k, v := range f {
			kvs[k] = v
		}
	}
	return kvs
}

func WithContext(ctx context.Context, a Adapter, fields ...types.M) Adapter {
	var kvs types.M
	l, ok := a.(*adapter)
	if !ok {
		kvs = merge(fields...)
		kvs = bindValues(ctx, kvs)
		return &adapter{adapter: a, fields: kvs, ctx: ctx}
	}
	kvs = types.M{}
	for k, v := range l.fields {
		kvs[k] = v
	}
	for k, v := range merge(fields...) {
		kvs[k] = v
	}
	kvs = bindValues(ctx, kvs)
	return &adapter{adapter: l.adapter, fields: kvs, ctx: ctx}
}

func With(a Adapter, fields types.M) Adapter {
	ctx := context.Background()
	return WithContext(ctx, a, fields)
}

// Valuer is returns a log value.
type Valuer func(ctx context.Context) interface{}

func bindValues(ctx context.Context, kvs types.M) types.M {
	_kvs := types.M{}
	for k, v := range kvs {
		if vf, ok := v.(Valuer); ok {
			_v := vf(ctx)
			if _v == nil {
				continue
			} else if _v == "" {
				_kvs[k] = "<empty>"
			} else {
				_kvs[k] = _v
			}
		}
	}
	return _kvs
}

func containsValuer(keyvals []interface{}) bool {
	for i := 1; i < len(keyvals); i += 2 {
		if _, ok := keyvals[i].(Valuer); ok {
			return true
		}
	}
	return false
}

func ValueFunc(key string) Valuer {
	return func(ctx context.Context) interface{} {
		return ctx.Value(key)
	}
}
