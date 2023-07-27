package config

import (
	"context"
)

type Store interface {
	Get(ctx context.Context) (Value, error)
}

type Retriever interface {
	Get(ctx context.Context, key string) (Value, error)
	Scan(ctx context.Context, key string, v interface{}) error
	Watch(ctx context.Context, key string) <-chan Change
	Close(ctx context.Context) error
}

type Change struct {
	Key      string
	NewValue Value
	OldValue Value
}

type Value interface {
	Int() (int64, error)
	Float() (float64, error)
	String() (string, error)
	Bool() (bool, error)
	Scan(v interface{}) error
	Map() (map[string]interface{}, error)
	Slice() ([]interface{}, error)
}
