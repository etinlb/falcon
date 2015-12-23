// Holds all the client handler functiosn
package main

import (
	"encoding/json"
	// "github.com/etinlb/falcon/core_lib"
	"github.com/etinlb/falcon/logger"
	"github.com/etinlb/falcon/network"
	"github.com/gorilla/websocket"
	"math/rand"
)

// Client events is data sent from the client to the server
func HandleClientEvent(event []byte, conn *websocket.Conn) *network.Message {
	logger.Trace.Printf("Received event: %s\n", string(event))

	// Putt in the input queue
	var message network.Message
	json.Unmarshal(event, &message)

	clientId := connections[conn]
	clientData := clientIdMap[clientId]
	logger.Trace.Printf("Queueing\n")
	clientData.QueueMessage(message)
	logger.Trace.Printf("Queued\n")
	return nil
	// return &message
}

func initializeClientData(conn *websocket.Conn) *network.Message {
	// initialize the connection
	logger.Info.Println("Connecting client")
	clientData := network.NewClient(conn)
	clientId := AddClientDataToMap(clientIdMap, &clientData)
	clientData.ClientId = clientId

	connectionMessage := MakeConnectionMessage(clientData)
	// make a add player event

	// clientData.QueueMessage()
	AddClientIdToMap(&clientData, clientId)
	// clientData.Player = AddNewPlayer(clientData)
	logger.Info.Printf("Syncing with %+v\n", connectionMessage)
	return &connectionMessage
}

func MakeConnectionMessage(clientData network.ClientData) network.Message {
	rawClientData, err := json.Marshal(clientData.GameObjects)
	rawJsonData := json.RawMessage(rawClientData)
	if err != nil {
		logger.Error.Panicf("Couldn't format data: %+v. Err: %s\n\n", clientData, err)
	}

	connectionMessage := network.Message{Event: "connect", Data: &rawJsonData}
	return connectionMessage
}

// func AddNewPlayer(conn ClientData) Player {
//     player := NewPlayer(0, 0, "testId")
//     client.GameObjects[player.Id] = &player
//     // Add to game world
//     AddPlayerObjectToWorld(player)
//     return player
// }

// func ExcludeClient(client *websocket.Conn) map[*websocket.Conn]bool {
//     // makes a map with only this one client to pass to sendPackets
//     connections := make(map[*websocket.Conn]bool)
//     connections[client] = true

//     return connections
// }

func AddClientIdToMap(client *network.ClientData, clientId int) {
	connections[client.Socket] = clientId
}

func AddClientDataToMap(mapToAdd map[int]*network.ClientData, clientToAdd *network.ClientData) int {
	x := rand.Int()
	for {
		if _, ok := mapToAdd[x]; !ok {
			mapToAdd[x] = clientToAdd
			return x
		}
		x = rand.Int()
	}
}
