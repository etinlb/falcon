package main

import (
	"encoding/json"
	"github.com/etinlb/falcon/core_lib"
	"github.com/etinlb/falcon/logger"
	"github.com/etinlb/falcon/network"
	"time"
)

// Spawns the game loop and returns the channels to comminucate with the game
// TODO: Currently that is just the move channels, maybe return the ticker channel?
// TODO: TODO: Make it return channel of channels
func StartGameLoop(socketHandler network.NetworkController) {
	// about 16 milliseconds for 60 fps a second
	gameTick := time.NewTicker(time.Millisecond * 16)
	server_tick_rate := 10 // Broadcast to clients every 10 tick

	// Init Physics
	physicsSpace := core_lib.NewPhysicsSpace(400, 400, 0.016)
	// TODO: How should the phsysics bodies be initialized? Game logic needs
	// them to set speeds while the physics engine needs the bodies to
	// simulate. They should be the same map but where does it come from?
	// Take the space one for now
	physicsComponents = physicsSpace.Bodys

	// // Physics runs at 50 fps
	// physicsTick := time.NewTicker(time.Millisecond * 20)
	// timeStep := (time.Millisecond * 2).Seconds()

	// TODO: Figure out buffering properly
	// moveChannel := make(chan *MoveMessage, 10)
	// addChannel := make(chan *AddMessage, 10)
	// broadcastAddChannel := make(chan *AddMessage, 10)

	// // actual Game Loop. TODO: Should this be a function call?
	go func() {
		current_tick := 0
		// Run the game loop forever.
		for range gameTick.C {
			updates := ReadAllInputMessages()

			if len(updates) != 0 {
				ApplyMessages(updates, gameObjects)
			}
			// TODO: Run Update on GameObjects

			// TODO: Physics?
			physicsSpace.TickPhysics(physicsSpace.Tick)

			// TODO: Have this done with a channel I think...
			// broadCastGameObjects()
			current_tick = (current_tick + 1) % server_tick_rate
			if current_tick == 0 {
				SendServerTickToClients(socketHandler)
			}

		}
	}()

	// // Start phyics loop, give it the movement channel and it's ticker
	// go PhysicsLoop(physicsTick, moveChannel, timeStep)

	logger.Info.Println("Started Game Loop")
}

// Function that sends the "server tick" to the clients
// should be run at a reasonable pace, i.e. .1 seconds
func SendServerTickToClients(networkController network.NetworkController) {
	// TODO: This should build the common update that will be sent to all clients
	// first.
	for clientId, clientData := range clientIdMap {
		// build the update message
		sequenceNum := clientData.CurrentSequnceNumber
		update := clientData.Player.BuildUpdateMessage()
		networkController.Send(update, clientData.ClientData)

		// add the sequence number
		logger.Trace.Printf("Sending: %s who is on sequence nubmer %d this data %s\n", clientId, sequenceNum, string(*update.Data))
	}
}

func ApplyMessages(messages []network.Message, gameObjects map[string]GameObject) {
	logger.Info.Printf("%+v", messages)

	for _, messages := range messages {
		if messages.Event == "move" {
			var updateMessage UpdateMessage
			_ = json.Unmarshal(*messages.Data, &updateMessage)
			logger.Info.Printf("Moving message %+v \n", updateMessage)
			gameObject := physicsComponents[updateMessage.Id]
			gameObject.Move(updateMessage.Velocity.X, updateMessage.Velocity.Y)
		}
	}
}

func ReadAllInputMessages() []network.Message {
	updates := make([]network.Message, 0)
	for clientId, clientData := range clientIdMap {

		// for clientId, clientData := range clientIdMap {
		// build the update message
		clientInputs, sequenceNum := clientData.ReadWholeQueue()
		logger.Trace.Printf("Inputes are %+v \n", clientInputs)
		updates = append(updates, clientInputs...)
		// apply all the updates to the client
		// add the sequence number
		logger.Trace.Printf("For client %s, on sequence number: %d \n", clientId, sequenceNum)
	}

	return updates
}

// func ReadMoveMessage(rawMoveMessage []byte) MoveMessage {
// 	var moveMessage MoveMessage
// 	json.Unmarshal(rawMoveMessage, &moveMessage)
// 	return moveMessage
// }

// func AddPhysicsComp(comp *PhysicsComponent, id string) {
//     physicsComponents[id] = comp
// }

// func AddObjectToConnectionObject(objId string, obj GameObject, client ClientData) {
//     client.GameObjects[objId] = obj
// }

// func AddPlayerObjectToWorld(player Player) {
//     playerObjects[player.Id] = &player
//     gameObjects[player.Id] = &player
//     AddPhysicsComp(player.PhysicsComp, player.Id)
// }

// // Physics loops listens from move requests and
// func PhysicsLoop(physicsTick *time.Ticker, moveChannel chan *MoveMessage, timeStep float64) {
//     frameSimulated := 0
//     for range physicsTick.C {
//         // Read any movement updates
//         select {
//         // Right now, a move request only comes in through player movement
//         case msg := <-moveChannel:
//             id := msg.Id
//             if physicsComp, ok := physicsComponents[id]; ok {
//                 //do something here
//                 physicsComp.Move(msg.X, msg.Y)
//             }
//         default:
//             // Move on to other things
//         }

//         TickPhysics(timeStep)
//         // TODO: Send this to a channel after reading an event so we can listen
//         // in and know exactly which tick the event was registered
//         frameSimulated++
//     }
// }

// // Ticks the physics engine once by time elapsed
// func TickPhysics(timeElapsed float64) {
//     for _, gameObj := range physicsComponents {
//         // Basic movement for now
//         gameObj.Location.X += gameObj.Velocity.X * timeElapsed
//         gameObj.Location.Y += gameObj.Velocity.Y * timeElapsed
//     }
// }
