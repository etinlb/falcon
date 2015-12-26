package core_lib

// Game Object is a struct with various components, components themselves
// aren't game objects
type GameObject interface {
	Update()
	// ReadMessage() // process data it gets from the client
	// BuildAddMessage() AddMessage
	// BuildUpdateMessage() UpdateMessage // process data it gets from the client
}

// Component for the bare minimum representation of a game object
// :id     - unique object id
// :source - source id of the client or server the object belongs to
type BaseGameObjData struct {
	Id string `json:"id"`
	// SourceId int    `json:"sourceId"`
}

func NewBaseGameObjData() BaseGameObjData {
	id := UniqueShortId()
	return BaseGameObjData{id}
}

func NewPlayer(x, y int) Player {
	physicsComponenet := NewPhysicsComponent(x, y)

	// TODO: Make a source id?
	baseData := NewBaseGameObjData()
	playerObject := Player{
		PhysicsComp:     &physicsComponenet,
		BaseGameObjData: baseData,
	}

	return playerObject
}
