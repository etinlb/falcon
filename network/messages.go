package network

import (
	"encoding/json"
	"github.com/etinlb/falcon/core_lib"
)

type Message struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"` // how data is parsed depends on the event
	// SequenceNum int             `json:"seq_num"`
}

type Events struct {
	Events []Message `json:"events"`
}

// send all game objects that are currently in the game object map to the
// client connected
func SyncClient(client *ClientData, gameObjects []core_lib.GameObject) {
	// TODO: Assess whether or not this is going to be to slow
	// syncData := Events{Events: make([]Message, 0)}

	// for _, obj := range gameObjects {
	// 	// var gameObjects map[string]GameObject
	// 	addMessage := obj.BuildAddMessage()
	// 	addBytes, _ := json.Marshal(addMessage)
	// 	message := Message{Event: "add", Data: addBytes, SequenceNum: 1}
	// 	syncData.Events = append(syncData.Events, message)
	// }
}
