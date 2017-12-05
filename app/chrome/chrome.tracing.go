package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Tracing - https://chromedevtools.github.io/devtools-protocol/tot/Tracing/
EXPERIMENTAL
*/
type Tracing struct{}

/*
End stops trace events collection.
*/
func (Tracing) End(socket *Socket) error {
	command := &protocol.Command{
		method: "Tracing.end",
		params: nil,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetCategories gets supported tracing categories.
*/
func (Tracing) GetCategories(socket *Socket, params *tracing.GetCategoriesParams) error {
	command := &protocol.Command{
		method: "Tracing.getCategories",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RecordClockSyncMarker records a clock sync marker in the trace.
*/
func (Tracing) RecordClockSyncMarker(socket *Socket, params *tracing.RecordClockSyncMarkerParams) error {
	command := &protocol.Command{
		method: "Tracing.recordClockSyncMarker",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
RequestMemoryDump requests a global memory dump.
*/
func (Tracing) RequestMemoryDump(socket *Socket, params *tracing.RequestMemoryDumpParams) error {
	command := &protocol.Command{
		method: "Tracing.requestMemoryDump",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Start starts trace events collection.
*/
func (Tracing) Start(socket *Socket, params *tracing.StartParams) error {
	command := &protocol.Command{
		method: "Tracing.start",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnBufferUsage adds a handler to the Tracing.bufferUsage event. Tracing.bufferUsage fires when a
buffer is used.
*/
func (Tracing) OnBufferUsage(socket *Socket, callback func(event *tracing.BufferUsageEvent)) error {
	handler := protocol.NewEventHandler(
		"Tracing.bufferUsage",
		func(name string, params []byte) {
			event := &tracing.BufferUsageEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnDataCollected adds a handler to the Tracing.dataCollected event. Tracing.dataCollected fires when
tracing is stopped, collected events will be sent as a sequence of dataCollected events followed by
tracingComplete event. Contains an bucket of collected trace events.
*/
func (Tracing) OnDataCollected(socket *Socket, callback func(event *tracing.DataCollectedEvent)) error {
	handler := protocol.NewEventHandler(
		"Tracing.dataCollected",
		func(name string, params []byte) {
			event := &tracing.BufferUsageEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnTracingComplete adds a handler to the Tracing.tracingComplete event. Tracing.tracingComplete fires
tracing is stopped and there is no trace buffers pending flush, all data were delivered via
dataCollected events.
*/
func (Tracing) OnTracingComplete(socket *Socket, callback func(event *tracing.TracingCompleteEvent)) error {
	handler := protocol.NewEventHandler(
		"Tracing.tracingComplete",
		func(name string, params []byte) {
			event := &tracing.TracingCompleteEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}