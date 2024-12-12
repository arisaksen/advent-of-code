package ecs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type DrawSystem struct {
	Registry *Registry
}

func (s *DrawSystem) Draw(screen *ebiten.Image) {
	for _, e := range s.Registry.Query(TransformType, DrawType) {
		transform := e.GetComponent(TransformType).(*TransformComponent)
		draw := e.GetComponent(DrawType).(*DrawComponent)
		options := &ebiten.DrawImageOptions{}
		point := draw.EbitenImage.Bounds().Size()
		options.GeoM.Scale(float64(draw.TileSize)/float64(point.X), float64(draw.TileSize)/float64(point.Y))
		options.GeoM.Translate(float64(transform.X), float64(transform.Y))
		screen.DrawImage(draw.EbitenImage, options)
	}
}
