package konfig

import (
	"context"
	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/weflux/fluxworks/config"
)

type Store struct {
	conf kconfig.Config
}

func (s *Store) Get(ctx context.Context) (config.Value, error) {
	//v := config.Value{}
	//if err := s.conf.Scan(&v); err != nil {
	//	return nil, err
	//}
	//return v, nil
	return nil, nil
}

func NewStore(path string) config.Store {
	conf := kconfig.New(
		kconfig.WithSource(
			file.NewSource(path),
		),
	)
	return &Store{
		conf: conf,
	}
}

type Retriever struct {
	//store *Store
	conf kconfig.Config
}

func (r *Retriever) Scan(ctx context.Context, key string, v interface{}) error {
	if key == "" {
		return r.conf.Scan(v)
	}
	c, err := r.Get(ctx, key)
	if err != nil {
		return err
	}
	return c.Scan(v)
}

func (r *Retriever) Get(ctx context.Context, key string) (config.Value, error) {
	c := r.conf.Value(key)
	return NewValue(c), nil
}

func (r *Retriever) Watch(ctx context.Context, key string) <-chan config.Change {
	ch := make(chan config.Change, 1)
	var observer kconfig.Observer = func(key string, value kconfig.Value) {
		ch <- config.Change{Key: key, NewValue: NewValue(value)}
	}
	if err := r.conf.Watch(key, observer); err != nil {
	}
	return ch
}

func (r *Retriever) Close(ctx context.Context) error {
	return r.conf.Close()
}

func NewRetriever(path string) config.Retriever {
	conf := kconfig.New(
		kconfig.WithSource(
			file.NewSource(path),
		),
	)
	if err := conf.Load(); err != nil {
		panic(err)
	}
	return &Retriever{conf: conf}

}

type Value struct {
	kconfig.Value
}

func (v *Value) Slice() ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewValue(v kconfig.Value) config.Value {
	return &Value{Value: v}
}

func (v *Value) Map() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	if err := v.Value.Scan(&m); err != nil {
		return nil, err
	}
	return m, nil
}
