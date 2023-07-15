package zaplog

import (
	"github.com/fluxstack/fluxworks/core/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZapLog_Log(t *testing.T) {
	zlog, err := NewZapLogger(Options{Production: true})
	assert.NoError(t, err)
	adapter := NewAdapter(zlog)
	//log := log2.Wrap(adapter, log2.Fields{})
	LOG := log.New(log.Wrap(adapter, log.Fields{}))
	LOG.Infow(log.Fields{
		"msg": "hello",
		"uid": 1234,
	})
	LOG3 := LOG.With(log.Field("task", 2))
	LOG3.Infow(log.Fields{
		"uid": "23455",
	})
}
