/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
	"encoding/json"
	"errors"
	"sync"

	log "github.com/Sirupsen/logrus"
)

/*
NewSocketCmd generates a new generic socket command
*/
func NewSocketCmd(method string, params *interface{}) *SocketCmd {
	cmd := new(SocketCmd)
	cmd.method = method
	cmd.params = params
	return cmd
}

/*
SocketCmdIface is the interface definition for a socket command
*/
type SocketCmdIface interface {
	/*
		Done handles the command result
	*/
	Done(result []byte, err error)
	/*
		Method returns the name of the command
	*/
	Method() string
	/*
		Params returns the command parameters to be used
	*/
	Params() interface{}
	/*
		Run executes the command on the specified socket
	*/
	Run(socket *Socket) error
}

/*
SocketCmd is a generic SocketCmdIface type
*/
type SocketCmd struct {
	err    error
	method string
	params *interface{}
	result SocketScreenshotResult
	wg     sync.WaitGroup
}

/*
Done is a SocketCmdIface implementation
*/
func (cmd *NewSocketCmd) Done(result []byte, err error) {
	if err == nil {
		err = json.Unmarshal(result, &cmd.result)
	}
	cmd.err = err
	cmd.wg.Done()
}

/*
Method is a SocketCmdIface implementation
*/
func (cmd *SocketScreenshotCmd) Method() string {
	return cmd.method
}

/*
Params is a SocketCmdIface implementation
*/
func (cmd *SocketScreenshotCmd) Params() (params interface{}) {
	return cmd.params
}

/*
Run is a SocketCmdIface implementation
*/
func (cmd *SocketScreenshotCmd) Run(socket *Socket) error {
	cmd.wg.Add(1)
	socket.SendCommand(cmd)
	cmd.wg.Wait()
	return cmd.err
}

/*
SocketPayload is a representation of a WebSocket JSON payload
*/
type SocketPayload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

/*
NewSocketPayload generates a new SocketPayload pointer
*/
func NewSocketPayload(id int, method string, params interface{}) *SocketPayload {
	payload := new(SocketPayload)
	payload.ID = id
	payload.Method = method
	payload.Params = params
	return payload
}

/*
SocketResponse represents a socket message
*/
type SocketResponse struct {
	Error  SocketError     `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}

/*
SendCommand sends a command to a connected socket

The socket's command mutex is locked, the command counter is incremented, the
payload is sent to the socket connection and the mutex is unlocked. The command
is stored using the counter value as it's Id. When the command is executed and
the socket responds, handleCmd() is executed to generate a response
*/
func (socket *Socket) SendCommand(command socketCmdIface) int {
	socket.cmdMutex.Lock()
	defer socket.cmdMutex.Unlock()

	socket.cmdID++
	payload := &SocketPayload{
		socket.cmdID,
		command.Method(),
		command.Params(),
	}
	tmp, _ := json.Marshal(payload)
	log.Debugf("Sending %#v", string(tmp))
	if err := socket.conn.WriteJSON(payload); err != nil {
		command.Done(nil, err)
	}
	socket.commands[payload.ID] = command

	return payload.ID
}

func (socket *Socket) handleCmd(response *SocketResponse) {
	var err error
	log.Infof("Command #%d responding", response.ID)
	socket.cmdMutex.Lock()
	defer socket.cmdMutex.Unlock()

	if command, ok := socket.commands[response.ID]; !ok {
		log.Warnf("Command %d not found: result=%s err=%s", response.ID, response.Result, response.Error.Message)

	} else {
		delete(socket.commands, response.ID)

		if "" != response.Error.Message {
			err = errors.New(response.Error.Message)
		}

		command.Done(response.Result, err)
	}
}
