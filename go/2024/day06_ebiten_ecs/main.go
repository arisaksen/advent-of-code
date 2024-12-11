package main

import (
	"aoc/golang/2024/day06_ebiten_ecs/ecs"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"strings"
)

var (

	//go:embed puzzle1_test.txt
	puzzle1Test string // 10 * 10 pixels

	//go:embed puzzle1.txt
	puzzle1 string // 130 * 130 pixels
)

const (
	screenWidth  = 1024
	screenHeight = 1024

	moveTimer = 1
)

type Game struct {
	timer           int
	drawSystem      *ecs.DrawSystem
	moveSystem      *ecs.MoveSystem
	collisionSystem *ecs.CollisionSystem
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSystem.Draw(screen)

	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\n", ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Update() error {
	g.moveSystem.Move()
	g.collisionSystem.CheckNextTile()

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	inputLines := strings.Split(puzzle1, "\n")
	var (
		_                  = len(inputLines)
		tileCountForHeight = len(inputLines)
		tileSize           = screenWidth / tileCountForHeight
	)
	log.Print(len(inputLines), len(inputLines[0]))

	registry := ecs.Registry{}
	for y, line := range inputLines {
		for x, char := range line {
			if char == '<' || char == '>' || char == '^' || char == 'v' {
				guard := registry.NewEntity()
				guard.AddComponent(&ecs.DrawComponent{Char: char})
				guard.AddComponent(&ecs.TransformComponent{Y: y * tileSize, X: x * tileSize, Size: tileSize})
				guard.AddComponent(&ecs.MoveComponent{MoveSpeed: 1 * tileSize, MoveTime: moveTimer})
				guard.AddComponent(&ecs.CollisionComponent{TileSize: tileSize, ScreenWidth: screenWidth, ScreenHeight: screenHeight, TileCounter: 1})
			} else if char == '#' {
				wall := registry.NewEntity()
				wall.AddComponent(&ecs.DrawComponent{Char: char})
				wall.AddComponent(&ecs.TransformComponent{Y: y * tileSize, X: x * tileSize, Size: tileSize})
				wall.AddComponent(&ecs.CollisionComponent{TileSize: tileSize})
			} else if char == '.' {
				floor := registry.NewEntity()
				floor.AddComponent(&ecs.DrawComponent{Char: char})
				floor.AddComponent(&ecs.TransformComponent{Y: y * tileSize, X: x * tileSize, Size: tileSize})
				floor.AddComponent(&ecs.CollisionComponent{TileSize: tileSize})
			}
		}
	}

	game := Game{
		drawSystem:      &ecs.DrawSystem{Registry: &registry},
		moveSystem:      &ecs.MoveSystem{Registry: &registry},
		collisionSystem: &ecs.CollisionSystem{Registry: &registry},
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("render single image")
	log.Fatal(ebiten.RunGame(&game))
}
