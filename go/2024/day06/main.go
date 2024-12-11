package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//go:embed puzzle1.txt
var puzzle1 string

//go:embed puzzle1_test.txt
var puzzle1Test string

const (
	screenWidth  = 1040
	screenHeight = screenWidth
)

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

var (
	inputLines         []string
	xGridCountInScreen int
	yGridCountInScreen int
	gridSize           float32
)

type Position struct {
	x int
	y int
}

type Game struct {
	moveDirection int
	gridChars     int
	//snakeBody     Position
	//apple         Position
	timer     int
	moveTime  int
	score     int
	bestScore int
}

func init() {
	rand.Seed(time.Now().UnixNano())
	inputLines = strings.Split(puzzle1Test, "\n")
	xGridCountInScreen = len(inputLines[0]) // 130
	yGridCountInScreen = xGridCountInScreen
	gridSize = float32(screenWidth / xGridCountInScreen)
}

func (g *Game) collidesWithWall() bool {
	//g.score = g.timer

	return g.snakeBody.x < 0 ||
		g.snakeBody.y < 0 ||
		g.snakeBody.x >= xGridCountInScreen ||
		g.snakeBody.y >= yGridCountInScreen
}

func (g *Game) needsToMoveSnake() bool {
	return g.timer%g.moveTime == 0
}

func (g *Game) reset() {
	fmt.Println("HERE")

	g.snakeBody.x = xGridCountInScreen / 2
	g.snakeBody.y = yGridCountInScreen / 2
	g.score = g.timer
	g.timer = 0
	//g.moveDirection = dirNone
}

func (g *Game) Update() error {
	//g.moveDirection = dirUp

	if g.needsToMoveSnake() {
		if g.collidesWithWall() {
			g.reset()
		}

		switch g.moveDirection {
		case dirLeft:
			g.snakeBody.x--
		case dirRight:
			g.snakeBody.x++
		case dirDown:
			g.snakeBody.y++
		case dirUp:
			g.snakeBody.y--
		}
	}

	g.timer++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	v := g.snakeBody
	vector.DrawFilledRect(screen, float32(v.x)*gridSize, float32(v.y)*gridSize, gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff}, false)
	for x, line := range inputLines {
		or
		y, char := range line{
			if char == '#'{
			vector.DrawFilledRect(screen, float32(xGridCountInScreen-x)*gridSize, float32(y)*gridSize, gridSize, gridSize, color.Black, false)
		}
		}
	}
	//vector.DrawFilledRect(screen, float32(g.apple.x*gridSize), float32(g.apple.y*gridSize), gridSize, gridSize, color.RGBA{0xFF, 0x00, 0x00, 0xff}, false)

	//if g.moveDirection == dirNone {
	//	ebitenutil.DebugPrint(screen, fmt.Sprintf("Press up/down/left/right to start"))
	//} else {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f Timer: %d, Score: %d", ebiten.ActualFPS(), g.timer, g.score))
	//}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func newGame() *Game {
	g := &Game{
		//apple:     Position{x: 3 * gridSize, y: 3 * gridSize},
		moveTime: 1,
		//snakeBody: make([]Position, 1),
	}
	g.snakeBody.x = xGridCountInScreen / 2
	g.snakeBody.y = yGridCountInScreen / 2
	g.moveDirection = dirUp
	return g
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Advent of Code - Day06")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
