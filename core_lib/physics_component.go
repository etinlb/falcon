package core_lib

var PlayerMovementXVel = 1000.0
var PlayerMovementYVel = 1000.0

type RectComponent struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

type PhysicsComponent struct {
    Location Vector2D `json:"location"`
    Velocity Vector2D `json:"velocity"`
    Force    Vector2D
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
