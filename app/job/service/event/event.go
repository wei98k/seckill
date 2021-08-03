package event

import "context"

type Message interface {
	Key() string
	Value() []byte
	Header() map[string]string
}

type Handler func(context.Context, Message) error


type Receiver interface {
	Receive(ctx context.Context, handler Handler) error
	Close() error
}
