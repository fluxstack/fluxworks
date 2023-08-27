package cluster

import (
	"github.com/weaveworks/mesh"
	"github.com/weflux/fluxworks/log"
)

type Swarm struct {
	mesh.Gossiper
	log *log.Logger
}
