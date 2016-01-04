// Holds all the client handler functiosn
package main

import (
	"encoding/json"
	"github.com/etinlb/falcon/core_lib"
	"github.com/etinlb/falcon/logger"
	"github.com/etinlb/falcon/network"
	"github.com/gorilla/websocket"
)

type Client struct {
	network.ClientData
	Player Player `json:"player"`
}

// Client events is data sent from the client to the server
func HandleClientEvent(event []byte, conn *websocket.Conn) *network.Message {
	logger.Trace.Printf("Received event: %s\n", string(event))

	// Putt in the input queue
	var message network.Message
	json.Unmarshal(event, &message)

	clientId := connections[conn]
	clientData := clientIdMap[clientId]
	logger.Trace.Printf("Queueing %+v data is %s\n", message, string(*message.Data))
	clientData.QueueMessage(message)
	logger.Trace.Printf("Queued %+v \n", clientData.InputQueue)
	return nil
	// return &message
}

func initializeClientData(conn *websocket.Conn) *network.Message {
	// initialize the connection
	logger.Info.Println("Connecting client")

	// Init all data that a client uses
	clientData := network.NewClient(conn)
	client := Client{ClientData: clientData}
	client.Player = NewPlayer(100, 300)

	// TODO: Add players to the game objects and physics component structs
	AddPhysicsComp(physicsComponents, client.Player.PhysicsComp, client.Player.Id)
	AddPlayerObjectToWorld(client.Player)

	logger.Info.Printf("%+v game Objects\n", gameObjects)

	AddClientDataToMap(clientIdMap, &client)
	AddClientToIdMap(&client, connections)

	// clientData.Player :=`
	connectionMessage := MakeConnectionMessage(client)
	// make a add player event

	// clientData.QueueMessage()
	// clientData.Player = AddNewPlayer(clientData)
	logger.Info.Printf("Syncing with %+v\n", connectionMessage)
	return &connectionMessage
}

// Makes the initial data message to send to the a connecting client.
func MakeConnectionMessage(client Client) network.Message {
	rawClientData, err := json.Marshal(client)
	rawJsonData := json.RawMessage(rawClientData)
	if err != nil {
		logger.Error.Panicf("Couldn't format data: %+v. Err: %s\n\n", client, err)
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

func AddClientToIdMap(client *Client, connections map[*websocket.Conn]string) {
	connections[client.Socket] = client.Id
}

func AddPlayerObjectToWorld(player Player) {
	gameObjects[player.Id] = &player
}

func AddPhysicsComp(physicsCompMap map[string]*core_lib.PhysicsComponent, comp *core_lib.PhysicsComponent, id string) {
	physicsCompMap[id] = comp
}

func AddClientDataToMap(mapToAdd map[string]*Client, clientToAdd *Client) {
	id := clientToAdd.Id
	mapToAdd[id] = clientToAdd

	// core_lib.UniqueShortId()
	// x := rand.Int()
	// for {
	// 	if _, ok := mapToAdd[x]; !ok {
	// 		mapToAdd[x] = clientToAdd
	// 		return x
	// 	}
	// 	x = rand.Int()
	// }
}
