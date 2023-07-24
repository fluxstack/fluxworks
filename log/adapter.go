package log

type Adapter interface {
	Log(level Level, fields Fields) error
}

type adapterWrap struct {
	fields  Fields
	adapter Adapter
}

func (w *adapterWrap) Log(level Level, fields Fields) error {
	kvs := make(map[string]interface{}, len(w.fields)+len(fields))
	for k, v := range w.fields {
		kvs[k] = v
	}
	for k, v := range fields {
		kvs[k] = v
	}
	return w.adapter.Log(level, kvs)
}

func With(logger Adapter, fields Fields) Adapter {
	l, ok := logger.(*adapterWrap)
	if !ok {
		return &adapterWrap{adapter: logger, fields: fields}
	}
	kvs := make(map[string]interface{}, len(l.fields)+len(fields))
	for k, v := range l.fields {
		kvs[k] = v
	}
	for k, v := range fields {
		kvs[k] = v
	}

	return &adapterWrap{adapter: l.adapter, fields: kvs}
}
