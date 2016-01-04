package network

import (
	"encoding/json"
	"github.com/etinlb/falcon/core_lib"
	"github.com/gorilla/websocket"
	"sync"
)

type NetworkedGameObjects interface {
	core_lib.GameObject
	BuildAddMessage() Message
}

type ClientConnectMessage struct {
	PlayerId string `json:"player_id"`
	Latency  int
}

// keeps track of data from a client
// Ignores Input queue and Socket when marshalling
type ClientData struct {
	Socket               *websocket.Conn `json:"-"`
	CurrentSequnceNumber int             `json:"sequenceNumber"`
	InputQueue           *MessageQueue   `json:"-"`
	Id                   string
}

type MessageQueue struct {
	sync.Mutex
	Queue []Message
}

type ClientMessage struct {
	Id   string `json:"id"`
	Data json.RawMessage
}

func NewClient(conn *websocket.Conn) (client ClientData) {
	messageQueue := MessageQueue{Queue: make([]Message, 0)}
	newClient := ClientData{}

	newClient.Socket = conn
	newClient.InputQueue = &messageQueue
	newClient.CurrentSequnceNumber = 0
	// TODO: Do we trust unique short id to be unique? Yeah probably.
	newClient.Id = core_lib.UniqueShortId()

	return newClient
}

func NewMessageQueue() *MessageQueue {
	queue := NewQueue()
	messageQueue := MessageQueue{Queue: queue}
	return &messageQueue
}

func NewQueue() []Message {
	return make([]Message, 0)
}

func (c *ClientData) ReadWholeQueue() ([]Message, int) {
	c.InputQueue.Lock()

	messages := make([]Message, len(c.InputQueue.Queue))
	messages = c.InputQueue.Queue[:]
	c.InputQueue.Queue = NewQueue()
	sequenceNumber := c.CurrentSequnceNumber
	c.InputQueue.Unlock()

	return messages, sequenceNumber
}

func (c *ClientData) QueueMessage(message Message) {
	c.InputQueue.Lock()

	c.InputQueue.Queue = append(c.InputQueue.Queue, message)
	c.CurrentSequnceNumber += 1
	c.InputQueue.Unlock()
}
