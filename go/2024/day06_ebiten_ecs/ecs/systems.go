package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type DrawSystem struct {
	Registry *Registry
}

func drawBox(screen *ebiten.Image, transform *TransformComponent, boxColor color.Color) {
	left := float32(transform.X)
	top := float32(transform.Y)
	widht := float32(transform.Size)
	height := float32(transform.Size)

	vector.DrawFilledRect(screen, left, top, widht, height, boxColor, true)
}

func (s *DrawSystem) Draw(screen *ebiten.Image) {
	for _, e := range s.Registry.Query(TransformType, DrawType) {
		transform := e.GetComponent(TransformType).(*TransformComponent)
		draw := e.GetComponent(DrawType).(*DrawComponent)

		char := draw.Char
		if char == '#' {
			drawBox(screen, transform, color.RGBA{R: 0, G: 0, B: 255, A: 255})
		} else if char == '.' {
			drawBox(screen, transform, color.White)
		}
	}

	// draw guard last
	for _, e := range s.Registry.Query(TransformType, DrawType) {
		transform := e.GetComponent(TransformType).(*TransformComponent)
		draw := e.GetComponent(DrawType).(*DrawComponent)

		char := draw.Char
		if char == '<' || char == '>' || char == '^' || char == 'v' {
			drawBox(screen, transform, color.RGBA{R: 255, G: 0, B: 0, A: 255})
			break
		}
	}
}

type MoveSystem struct {
	Registry *Registry
}

func (m *MoveSystem) Move() {
	for _, e := range m.Registry.Query(MoveType, TransformType, DrawType) {
		transform := e.GetComponent(TransformType).(*TransformComponent)
		move := e.GetComponent(MoveType).(*MoveComponent)
		draw := e.GetComponent(DrawType).(*DrawComponent)

		if move.MoveTimer%move.MoveTime == 0 {
			if draw.Char == '<' {
				transform.X -= move.MoveSpeed
			} else if draw.Char == '>' {
				transform.X += move.MoveSpeed
			} else if draw.Char == '^' {
				transform.Y -= move.MoveSpeed
			} else if draw.Char == 'v' {
				transform.Y += move.MoveSpeed
			} else {
				log.Printf("Char '%c' does not support move", draw.Char)
			}
		}
		move.MoveTimer++
	}
}

type CollisionSystem struct {
	Registry *Registry
}

func (c *CollisionSystem) CheckNextTile() {
	for _, e := range c.Registry.Query(MoveType, TransformType, DrawType, CollisionType) {
		move := e.GetComponent(MoveType).(*MoveComponent)
		transform := e.GetComponent(TransformType).(*TransformComponent)
		draw := e.GetComponent(DrawType).(*DrawComponent)
		collision := e.GetComponent(CollisionType).(*CollisionComponent)

		if move.MoveTimer%move.MoveTime == 0 {
			if transform.X < 0 || transform.X > collision.ScreenWidth || transform.Y < 0 || transform.Y > collision.ScreenHeight {
				log.Printf("Outside Screen. TileCounter: %d", collision.TileCounter)
				continue
			}

			for _, otherEntity := range c.Registry.entities {
				drawOther := otherEntity.GetComponent(DrawType).(*DrawComponent)
				transformOther := otherEntity.GetComponent(TransformType).(*TransformComponent)
				if drawOther.Char == '.' && transformOther.X == transform.X && transformOther.Y == transform.Y {
					collision.TileCounter++
					drawOther.Char = '0'
				}
				if drawOther.Char != '#' {
					continue
				}
				if draw.Char == '<' {
					if transformOther.Y == transform.Y && transformOther.X == transform.X-collision.TileSize {
						log.Printf("collision detected at: %v", transformOther)
						draw.Char = '^'
					}
				} else if draw.Char == '>' {
					if transformOther.Y == transform.Y && transformOther.X == transform.X+collision.TileSize {
						log.Printf("collision detected at: %v", transformOther)
						draw.Char = 'v'
					}
				} else if draw.Char == '^' {
					if transformOther.X == transform.X && transformOther.Y == transform.Y-collision.TileSize {
						log.Printf("collision detected at: %d %d", transformOther.X, transformOther.Y)
						draw.Char = '>'
					}
				} else if draw.Char == 'v' {
					if transformOther.X == transform.X && transformOther.Y == transform.Y+collision.TileSize {
						log.Printf("collision detected at: %v", transformOther)
						draw.Char = '<'
					}
				}
			}
		}
	}
}
