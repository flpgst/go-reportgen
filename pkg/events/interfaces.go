package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispacherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispach(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handle EventHandlerInterface) bool
	Clear() error
}
