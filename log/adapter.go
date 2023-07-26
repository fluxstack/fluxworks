package log

import (
	"context"
	"github.com/samber/lo"
)

type Adapter interface {
	Log(level Level, fields Fields) error
}

type adapter struct {
	fields  Fields
	adapter Adapter
	prefix  []interface{}
	ctx     context.Context
}

func (w *adapter) Log(level Level, fields Fields) error {
	kvs := make(map[string]Field, len(w.fields)+len(fields))
	for _, f := range w.fields {
		kvs[f.K] = f
	}
	for _, f := range fields {
		kvs[f.K] = f
	}
	return w.adapter.Log(level, lo.Values(kvs))
}

func WithContext(ctx context.Context, adapter Adapter) Adapter {
	return nil
}

func With(a Adapter, fields Fields) Adapter {
	l, ok := a.(*adapter)
	if !ok {
		return &adapter{adapter: a, fields: fields}
	}
	kvs := make(map[string]Field, len(l.fields)+len(fields))
	for _, f := range l.fields {
		kvs[f.K] = f
	}
	for _, f := range fields {
		kvs[f.K] = f
	}

	return &adapter{adapter: l.adapter, fields: lo.Values(kvs)}
}

// Valuer is returns a log value.
type Valuer func(ctx context.Context) interface{}

func bindValues(ctx context.Context, keyvals []interface{}) {
	for i := 1; i < len(keyvals); i += 2 {
		if v, ok := keyvals[i].(Valuer); ok {
			keyvals[i] = v(ctx)
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
