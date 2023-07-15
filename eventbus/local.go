package eventbus

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/fluxstack/fluxworks/log"
)

type EventBus struct {
	message.Publisher
	message.Subscriber
	router *message.Router
	ctx    context.Context
}

type Listener message.NoPublishHandlerFunc
type Processor message.HandlerFunc

func New(logger *log.Logger) *EventBus {
	busLogger := NewLogger(logger)
	router, err := message.NewRouter(message.RouterConfig{}, busLogger)
	if err != nil {
		logger.Fatal("new eventbus error", err)
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
	bus.router.AddHandler(name, inTopic, bus, outTopic, bus, message.HandlerFunc(processor))
	if bus.router.IsRunning() {
		return bus.router.RunHandlers(bus.ctx)
	}
	return nil
}

func (bus *EventBus) IsRunning() bool {
	return bus.router.IsRunning()
}
