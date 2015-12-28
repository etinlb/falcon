package core_lib

var PlayerMovementXVel = 10.0
var PlayerMovementYVel = 10.0

type Vector2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// A physics space, with dimensions, gravity and physics components
type PhysicsSpace struct {
	Dimensions Vector2D
	Tick       float64
	Gravity    Vector2D
	Bodys      map[string]*PhysicsComponent
	// Bodys map[string]*PhysicsComponent
}

// Ticks the physics engine once by time elapsed
func (p *PhysicsSpace) TickPhysics(timeElapsed float64) {
	for _, gameObj := range p.Bodys {
		// Basic movement for now
		gameObj.Location.X += gameObj.Velocity.X * timeElapsed
		gameObj.Location.Y += gameObj.Velocity.Y * timeElapsed
		gameObj.Velocity.Y += p.Gravity.Y * timeElapsed

		if 0 >= (gameObj.Location.Y - gameObj.Dimensions.Y) {
			gameObj.Velocity.Y = 0 // hit the floor stop moving
			gameObj.Location.Y = gameObj.Dimensions.Y
		}
	}
}

// Ticks the physics engine once by time elapsed
func NewPhysicsSpace(width, height, tick float64) PhysicsSpace {
	space := PhysicsSpace{}
	dimensions := NewVector(width, height)

	space.Dimensions = dimensions
	space.Tick = tick

	space.Bodys = make(map[string]*PhysicsComponent)
	space.Gravity = NewVector(0, -9.8)

	return space
}

type PhysicsComponent struct {
	Location   Vector2D `json:"location"`
	Velocity   Vector2D `json:"velocity"`
	Force      Vector2D
	Dimensions Vector2D
}

func NewVector(x, y float64) Vector2D {
	rect := Vector2D{X: x, Y: y}

	return rect
}

func (m *PhysicsComponent) Move(xAxis, yAxis float64) {
	m.Velocity.X += PlayerMovementXVel * xAxis
	m.Velocity.Y += PlayerMovementYVel * yAxis
}

func NewPhysicsComponent(x, y float64) PhysicsComponent {
	locationVector := NewVector(x, y)
	gameObject := PhysicsComponent{Location: locationVector}

	return gameObject
}
