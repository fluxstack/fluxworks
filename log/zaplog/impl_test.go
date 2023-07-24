package zaplog

import (
	log2 "github.com/fluxstack/fluxworks/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZapLog_Log(t *testing.T) {
	zlog, err := NewZapLogger(Options{Production: true})
	assert.NoError(t, err)
	adapter := NewAdapter(zlog)
	//log := log2.With(adapter, log2.Fields{})
	LOG := log2.New(log2.With(adapter, log2.Fields{}))
	LOG.Infow(log2.Fields{
		"msg": "hello",
		"uid": 1234,
	})
	LOG3 := LOG.With(log2.Field("task", 2))
	LOG3.Infow(log2.Fields{
		"uid": "23455",
	})
}
