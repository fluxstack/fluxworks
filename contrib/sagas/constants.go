package sagas

// cqrs message headers
const (
	MessageEventPrefix     = "event_"
	MessageEventName       = MessageEventPrefix + "name"
	MessageEventEntityName = MessageEventPrefix + "entity_name"
	MessageEventEntityID   = MessageEventPrefix + "entity_id"

	MessageCommandPrefix       = "command_"
	MessageCommandName         = MessageCommandPrefix + "name"
	MessageCommandChannel      = MessageCommandPrefix + "channel"
	MessageCommandReplyChannel = MessageCommandPrefix + "reply_channel"

	MessageReplyPrefix  = "reply_"
	MessageReplyName    = MessageReplyPrefix + "name"
	MessageReplyOutcome = MessageReplyPrefix + "outcome"
)

const (
	notCompensating = false
	isCompensating  = true
)

// LifecycleHook type for hooking in custom code at specific stages of a saga
type LifecycleHook int

// Definition lifecycle hooks
const (
	SagaStarting LifecycleHook = iota
	SagaCompleted
	SagaCompensated
)

// Saga message headers
const (
	MessageCommandSagaID   = MessageCommandPrefix + "saga_id"
	MessageCommandSagaName = MessageCommandPrefix + "saga_name"
	MessageCommandResource = MessageCommandPrefix + "resource"

	MessageReplySagaID   = MessageReplyPrefix + "saga_id"
	MessageReplySagaName = MessageReplyPrefix + "saga_name"
)
