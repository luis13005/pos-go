package events

import (
	"sync"
	"time"
)

type Event interface {
	GetName() string
	GetDateTime() time.Time
	GetPayLoad() interface{}
}

type EventHandler interface {
	Handle(event Event, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandler) error
	Dispatch(event Event) error
	Remove(eventName string, handler EventHandler) error
	Has(eventName string, handler EventHandler) bool
	Clear() error
}
