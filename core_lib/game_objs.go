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

type BaseRectData struct {
	Velocity Vector2D
	Location Vector2D
}

func NewBaseGameObjData() BaseGameObjData {
	id := UniqueShortId()
	return BaseGameObjData{id}
}
