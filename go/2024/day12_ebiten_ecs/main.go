package main

import (
	"aoc/golang/2024/day12_ebiten_ecs/ecs"
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"path"
	"strings"
)

var (

	//go:embed puzzle1.txt
	puzzle1     string // 140 * 140 pixels
	puzzleTest1 = `AAAA
BBCD
BBCC
EEEC`
	puzzleTest2 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	//go:embed assets
	Assets embed.FS
)

const (
	screenWidth  = 1200
	screenHeight = 1200

	moveTimer = 1
)

type Game struct {
	timer      int
	drawSystem *ecs.DrawSystem
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawSystem.Draw(screen)

	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\n", ebiten.ActualTPS(), ebiten.ActualFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	//input := puzzleTest1
	input := puzzleTest2
	//input := puzzle1

	inputLines := strings.Split(input, "\n")
	var (
		mapBoarderTile     = 1
		_                  = len(inputLines) + 2*mapBoarderTile
		tileCountForHeight = len(inputLines) + 2*mapBoarderTile
		tileSize           = screenWidth / tileCountForHeight
	)
	log.Print(len(inputLines), len(inputLines[0]))

	paths, err := GetAllBinFilenames(&Assets, "assets")
	if err != nil {
		log.Fatal(err)
	}
	counter := len(paths) - 1
	vegetables := make(map[rune]string)
	for _, char := range input {
		_, ok := vegetables[char]
		if !ok {
			p := paths[counter]
			vegetables[char] = p
			counter--
			if counter == 0 {
				counter = len(paths) - 1
			}
		}
	}

	registry := ecs.Registry{}
	for y, line := range inputLines {
		for x, char := range line {
			e := registry.NewEntity()
			e.AddComponent(&ecs.TransformComponent{Y: y*tileSize + mapBoarderTile*tileSize, X: x*tileSize + mapBoarderTile*tileSize, Size: tileSize})
			fileReader, err := Assets.Open(vegetables[char])
			if err != nil {
				log.Fatal(err, char, vegetables)
			}
			ebitenImage, _, err := ebitenutil.NewImageFromReader(fileReader)
			if err != nil {
				log.Fatal(err)
			}
			e.AddComponent(&ecs.DrawComponent{Char: char, EbitenImage: ebitenImage, TileSize: tileSize})
		}
	}

	game := Game{
		drawSystem: &ecs.DrawSystem{Registry: &registry},
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("render single image")
	log.Fatal(ebiten.RunGame(&game))
}

func GetAllBinFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := GetAllBinFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}
