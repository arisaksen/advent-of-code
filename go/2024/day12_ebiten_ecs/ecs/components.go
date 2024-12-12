package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ComponentType defines the supported component types in a user readable format
type ComponentType string

const (
	TransformType ComponentType = "TRANSFORM"
	DrawType      ComponentType = "DRAW"
)

// ComponentTyper returns the type of component
type ComponentTyper interface{ Type() ComponentType }

type TransformComponent struct {
	X, Y int
	Size int
}

func (t *TransformComponent) Type() ComponentType { return TransformType }

type DrawComponent struct {
	Char        rune
	EbitenImage *ebiten.Image
	TileSize    int
}

func (t *DrawComponent) Type() ComponentType { return DrawType }
