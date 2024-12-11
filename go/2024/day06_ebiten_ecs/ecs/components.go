package ecs

// ComponentType defines the supported component types in a user readable format
type ComponentType string

const (
	DrawType      ComponentType = "DRAW"
	TransformType ComponentType = "TRANSFORM"
	CollisionType ComponentType = "COLLISION"
	MoveType      ComponentType = "MOVE"
)

// ComponentTyper returns the type of component
type ComponentTyper interface{ Type() ComponentType }

type TransformComponent struct {
	X, Y int
	Size int
}

func (t *TransformComponent) Type() ComponentType { return TransformType }

type CollisionComponent struct {
	TileSize                  int
	ScreenWidth, ScreenHeight int
	TileCounter               int
}

func (t *CollisionComponent) Type() ComponentType { return CollisionType }

type MoveComponent struct {
	MoveTimer   int
	MoveSpeed   int
	MoveTime    int
	MoveCounter int
}

func (t *MoveComponent) Type() ComponentType { return MoveType }

type DrawComponent struct {
	Char rune
}

func (t *DrawComponent) Type() ComponentType { return DrawType }
