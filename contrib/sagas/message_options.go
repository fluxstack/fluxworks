package sagas

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

// WithSagaInfo is an option to set additional Saga specific headers
func WithSagaInfo(instance *Instance) MessageOption {
	return WithHeaders(map[string]string{
		MessageCommandSagaID:   instance.sagaID,
		MessageCommandSagaName: instance.sagaName,
	})
}

// WithDestinationChannel is and option to set the destination of the outgoing Message
//
// This will override the previous value set by interface { DestinationChannel() string }
func WithDestinationChannel(destinationChannel string) MessageOption {
	return func(m *message.Message) {
		m.Metadata.Set(MessageChannel, destinationChannel)
	}
}

const (
	MessageChannel = "channel"
)
