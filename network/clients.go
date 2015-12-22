package network

import (
	"encoding/json"
	"github.com/etinlb/falcon/core_lib"
	"github.com/gorilla/websocket"
	"sync"
	// "sync/atomic"
)

type ClientConnectMessage struct {
	PlayerId string `json:"player_id"`
	Latency  int
}

// keeps track of data from a client
type ClientData struct {
	Socket               *websocket.Conn
	GameObjects          map[string]core_lib.GameObject
	ClientId             int
	CurrentSequnceNumber int
	InputQueue           *MessageQueue
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
	messageQueue := MessageQueue{Queue: make([]Message, 1)}
	newClient := ClientData{Socket: conn, GameObjects: make(map[string]core_lib.GameObject), InputQueue: &messageQueue}
	newClient.CurrentSequnceNumber = 0
	return newClient
}

func NewMessageQueue() *MessageQueue {
	queue := NewQueue()
	messageQueue := MessageQueue{Queue: queue}
	return &messageQueue
}

func NewQueue() []Message {
	return make([]Message, 1)
}

func (c *ClientData) ReadWholeQueue() ([]Message, int) {
	c.InputQueue.Lock()

	messages := make([]Message, len(c.InputQueue.Queue))
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
