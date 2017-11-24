/*
Package chrome provides an interface to a headless Chrome instance.
*/
package chrome

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

/*
Socket represents a websocket connection to the Browser instance
*/
type Socket struct {
	cmdID    int
	cmdMutex sync.Mutex
	evtMutex sync.Mutex
	events   map[string][]EventInterface
	commands map[int]SocketCmdIface // key is id.
	conn     *websocket.Conn
}

/*
NewSocket returns a new websocket connection
*/
func NewSocket(tab *Tab) (*Socket, error) {

	dialer := &websocket.Dialer{
		EnableCompression: false,
	}
	header := http.Header{
		"Origin": []string{tab.WebSocketDebuggerURL},
	}

	webSocket, response, err := dialer.Dial(tab.WebSocketDebuggerURL, header)
	if err != nil {
		log.Warningf("Could not create socket connection. %s responded with '%s'", tab.WebSocketDebuggerURL, response.Status)
		return nil, err
	}

	socket := &Socket{
		conn:     webSocket,
		commands: make(map[int]socketCmdIface),
		events:   make(map[string][]EventInterface),
	}
	go socket.Listen(tab)

	log.Infof("New socket connection listening on %s", tab.WebSocketDebuggerURL)
	return socket, nil
}

/*
Close closes the current socket connection
*/
func (socket *Socket) Close() error {
	return socket.conn.Close()
}

/*
Listen starts the socket read loop
*/
func (socket *Socket) Listen(tab *Tab) {
	for {
		response := &SocketResponse{}
		err := socket.conn.ReadJSON(response)
		if err != nil {
			log.Error(err)
			if err == io.EOF ||
				websocket.IsCloseError(err, 1006) ||
				strings.Contains(err.Error(), "use of closed network connection") {
				log.Error(err)
				break
			}
		} else if response.ID > 0 {
			socket.handleCmd(response)
		} else {
			params, _ := json.Marshal(response.Params)
			log.Infof("%s: %s", tab.ID, response.Method)
			log.Debugf("Event params: %s", params)
			socket.handleEvent(response)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
