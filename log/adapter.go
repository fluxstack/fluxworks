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

func WithContext(ctx context.Context, a Adapter, fields types.M) Adapter {
	var kvs types.M
	l, ok := a.(*adapter)
	if !ok {
		kvs = fields
		bindValues(ctx, kvs)
		return &adapter{adapter: a, fields: kvs, ctx: ctx}
	}
	kvs = make(types.M, len(l.fields)+len(fields))
	for k, v := range l.fields {
		kvs[k] = v
	}
	bindValues(ctx, kvs)
	return &adapter{adapter: l.adapter, fields: kvs, ctx: ctx}
}

func With(a Adapter, fields types.M) Adapter {
	l, ok := a.(*adapter)
	if !ok {
		return &adapter{adapter: a, fields: fields}
	}
	kvs := make(types.M, len(l.fields)+len(fields))
	for k, v := range l.fields {
		kvs[k] = v
	}
	for k, v := range fields {
		kvs[k] = v
	}

	return &adapter{adapter: l.adapter, fields: kvs}
}

// Valuer is returns a log value.
type Valuer func(ctx context.Context) interface{}

func bindValues(ctx context.Context, kvs types.M) {
	for k, v := range kvs {
		if vf, ok := v.(Valuer); ok {
			kvs[k] = vf(ctx)
		}
	}
}

func containsValuer(keyvals []interface{}) bool {
	for i := 1; i < len(keyvals); i += 2 {
		if _, ok := keyvals[i].(Valuer); ok {
			return true
		}
	}
	return false
}
