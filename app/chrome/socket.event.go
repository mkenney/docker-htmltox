/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
	log "github.com/Sirupsen/logrus"
)

/*
EventInterface is an interface
*/
type EventInterface interface {
	OnEvent(name string, params []byte)
}
type simpleEvent struct {
	cb func(name string, params []byte)
}

/*
OnEvent
*/
func (s *simpleEvent) OnEvent(name string, params []byte) {
	s.cb(name, params)
}

func FuncToEvent(fn func(name string, params []byte)) EventInterface {
	return &simpleEvent{fn}
}

/*
AddEvent adds an event to the event stack
*/
func (socket *Socket) AddEventHandler(name string, cb func(name string, params []byte)) {
	handler := FuncToEvent(cb)
	socket.evtMutex.Lock()
	defer socket.evtMutex.Unlock()
	for _, registeredEvent := range socket.events[name] {
		if registeredEvent == handler {
			return
		}
	}
	socket.events[name] = append(socket.events[name], handler)
}

/*
RemoveEvent removes an event from the stack
*/
func (socket *Socket) RemoveEvent(name string, event EventInterface) {
	socket.evtMutex.Lock()
	defer socket.evtMutex.Unlock()
	events := socket.events[name]
	for i, s := range events {
		if s == event {
			l := len(events)
			events[i] = events[l-1]
			socket.events[name] = events[:l-1]
			return
		}
	}
}

/*
NewEvent creates a new event struct
*/
func NewEvent(cb func(name string, params []byte)) EventInterface {
	return &simpleEvent{cb}
}

func (socket *Socket) handleEvent(response *SocketResponse) {
	if response.Method == "Inspector.targetCrashed" {
		log.Fatalf("Chrome has crashed!")
	}
	socket.evtMutex.Lock()
	defer socket.evtMutex.Unlock()

	log.Debugf("Event received: '%v'", response)
	for _, event := range socket.events[response.Method] {
		log.Debugf("%s handler(s) found", response.Method)
		go event.OnEvent(response.Method, []byte(response.Params))
	}
}
