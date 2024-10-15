package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	tileImage *ebiten.Image
	board     *Board
}

var backgroundColor = color.Gray{0xCC}
var tileColor = color.Black

const tileSize = 20
const screenWidth = 640
const screenHeight = 480

func NewGame(boardWidth int, boardHeight int) (*Game, error) {
	// todo: there should be a NewBoard
	// todo: robot and robot coordinate is dead code / bufd - big upfront design
	board := &Board{
		Coordinate{1, 1},
		boardWidth,
		boardHeight,
	}
	game := &Game{
		tileImage: nil,
		board:     board,
	}

	game.tileImage = ebiten.NewImage(tileSize, tileSize)
	game.tileImage.Fill(tileColor)

	// todo: error handling is not used - yagni, bufd
	return game, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// todo: window size, resolution, tileSize and tileOffset should'n be hard coded (full screen, use maths to calculate)
	tileOffset := tileSize + 2

	// todo: there should be a board.draw where the tileImage is passed
	screen.Fill(backgroundColor)

	for row := 0; row < g.board.height; row++ {
		for column := 0; column < g.board.width; column++ {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(column*tileOffset), float64(row*tileOffset))
			screen.DrawImage(g.tileImage, options)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth / 2, screenHeight / 2
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)

	ebiten.SetWindowTitle("smcd go labyrinth")

	// todo: width and height shouldn't be hard coded
	game, err := NewGame(2, 2)
	if err = ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
