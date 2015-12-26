package core_lib

var PlayerMovementXVel = 10
var PlayerMovementYVel = 10

type Vector2D struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PhysicsComponent struct {
	Location Vector2D `json:"location"`
	Velocity Vector2D `json:"velocity"`
	Force    Vector2D
}

func NewVector(x, y int) Vector2D {
	rect := Vector2D{X: x, Y: y}

	return rect
}

func (m *PhysicsComponent) Move(xAxis, yAxis int) {
	m.Velocity.X += PlayerMovementXVel * xAxis
	m.Velocity.Y += PlayerMovementYVel * yAxis
}

func NewPhysicsComponent(x, y int) PhysicsComponent {
	locationVector := NewVector(x, y)
	gameObject := PhysicsComponent{Location: locationVector}

	return gameObject
}

// ==============TODO: Evaluate where player should be. ==============
type Player struct {
	BaseGameObjData
	PhysicsComp *PhysicsComponent
}

func (p *Player) Move(xAxis, yAxis int) {
	p.PhysicsComp.Move(xAxis, yAxis)
}
