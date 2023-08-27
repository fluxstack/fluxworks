package cluster

import (
	"github.com/weaveworks/mesh"
	"github.com/weflux/fluxworks/logging"
)

type Swarm struct {
	mesh.Gossiper
	log *logging.Logger
}
