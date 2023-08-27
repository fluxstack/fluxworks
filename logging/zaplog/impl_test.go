package zaplog

import (
	"github.com/stretchr/testify/assert"
	"github.com/weflux/fluxworks/logging"
	"github.com/weflux/fluxworks/types"
	"testing"
)

func TestZapLog_Log(t *testing.T) {
	zlog, err := NewZapLogger(Options{Production: true})
	assert.NoError(t, err)
	adapter := NewAdapter(zlog)
	//log := logging.With(adapter, logging.Fields{})
	LOG := logging.New(logging.With(adapter, types.M{}))
	LOG.Info("hello",
		types.M{
			"uid": 1234,
		},
	)
	LOG3 := LOG.With(types.M{"task": 2})

	LOG3.Info("", types.M{
		"uid": "23455",
	})
}
