package sagas

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

// JUST FOR TEST
type memStore struct {
	mu             sync.RWMutex
	instancesByIds map[string]*Instance
}

func NewMemStore() InstanceStore {
	return &memStore{instancesByIds: map[string]*Instance{}}
}

func (ms *memStore) Find(ctx context.Context, sagaID string, data SagaData) (*Instance, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	ins, ok := ms.instancesByIds[sagaID]
	if !ok {
		return nil, errors.New("instance not found")
	}
	return ins, nil
}

func (ms *memStore) Save(ctx context.Context, sagaInstance *Instance) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.instancesByIds[sagaInstance.sagaID] = sagaInstance
	return nil
}

func (ms *memStore) Update(ctx context.Context, sagaInstance *Instance) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.instancesByIds[sagaInstance.sagaID] = sagaInstance
	return nil
}
