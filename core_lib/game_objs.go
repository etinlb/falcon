package core_lib

// var gameId = 1
// var playerMovementXVel = 1000.0
// var playerMovementYVel = 1000.0

// Game Object is a struct with various components, components themselves
// aren't game objects
type GameObject interface {
	Update()
	// ReadMessage() // process data it gets from the client
	// BuildAddMessage() AddMessage
	// BuildUpdateMessage() UpdateMessage // process data it gets from the client
}