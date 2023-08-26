package messaging

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/fluxstack/fluxworks/log"
	"github.com/golang/protobuf/proto"
)

type EventBus struct {
	message.Publisher
	message.Subscriber
	router *message.Router
	ctx    context.Context
}

type Listener message.NoPublishHandlerFunc
type Processor message.HandlerFunc

func NewLocal(logger *log.Logger) *EventBus {
	busLogger := NewLogger(logger)
	router, err := message.NewRouter(message.RouterConfig{}, busLogger)
	if err != nil {
		logger.Error("new eventbus error", err)
		return nil
	}
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, busLogger)

	return &EventBus{
		Publisher:  pubSub,
		Subscriber: pubSub,
		router:     router,
		ctx:        context.Background(),
	}
}

func (bus *EventBus) Start(ctx context.Context) error {
	return bus.router.Run(ctx)
}

func (bus *EventBus) AddListener(name, topic string, listener Listener) error {
	bus.router.AddNoPublisherHandler(name, topic, bus, message.NoPublishHandlerFunc(listener))
	if bus.router.IsRunning() {
		return bus.router.RunHandlers(bus.ctx)
	}
	return nil
}

func (bus *EventBus) AddProcessor(name, inTopic string, outTopic string, processor Processor) error {
	bus.router.AddHandler(name, inTopic, bus, outTopic, bus.Publisher, message.HandlerFunc(processor))
	if bus.router.IsRunning() {
		return bus.router.RunHandlers(bus.ctx)
	}
	return nil
}

func (bus *EventBus) IsRunning() bool {
	return bus.router.IsRunning()
}

type Envelope interface {
	ToMessage() (*message.Message, error)
	MustToMessage() *message.Message
}

type JsonEnvelope struct {
	Message interface{}
}

func (ev *JsonEnvelope) ToMessage() (*message.Message, error) {
	bs, err := json.Marshal(ev.Message)
	if err != nil {
		return nil, err
	}
	return message.NewMessage(watermill.NewUUID(), bs), nil
}

func (ev *JsonEnvelope) MustToMessage() *message.Message {
	bs, err := json.Marshal(ev.Message)
	if err != nil {
		panic(err)
	}
	return message.NewMessage(watermill.NewUUID(), bs)
}

type ProtoEnvelope struct {
	Message proto.Message
}

func (ev *ProtoEnvelope) ToMessage() (*message.Message, error) {
	bs, err := proto.Marshal(ev.Message)
	if err != nil {
		return nil, err
	}
	return message.NewMessage(watermill.NewUUID(), bs), nil
}

func (ev *ProtoEnvelope) MustToMessage() *message.Message {
	bs, err := proto.Marshal(ev.Message)
	if err != nil {
		panic(err)
	}
	return message.NewMessage(watermill.NewUUID(), bs)
}

func (bus *EventBus) Publish(topic string, envelope Envelope) error {
	return bus.Publisher.Publish(topic, envelope.MustToMessage())
}

func (bus *EventBus) Request(topic string, req Envelope) (Envelope, error) {
	return nil, nil
}
