package events

import (
	"errors"
	"sync"
)

var errHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return errHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if _, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, h := range ed.handlers[event.GetName()] {
			wg.Add(1)
			h.Handle(event, wg)
		}
	}
	return nil
}

func (ed *EventDispatcher) Remove(event string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[event]; ok {
		for i, h := range ed.handlers[event] {
			if h == handler {
				ed.handlers[event] = append(ed.handlers[event][:i], ed.handlers[event][i+1:]...)
			}
		}
	}

	return nil
}
