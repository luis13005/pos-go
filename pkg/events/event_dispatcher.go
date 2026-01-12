package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("Handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandler) error {

	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)

	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandler)

}
