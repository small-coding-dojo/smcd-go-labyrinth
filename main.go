package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct {
	tileImage *ebiten.Image
	board     *Board
}

func NewGame() (*Game, error) {
	// todo: there should be a NewBoard
	board := &Board{
		Coordinate{1, 1},
		2,
		2,
	}
	game := &Game{
		tileImage: nil,
		board:     board,
	}
	return game, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tileSize := 20
	tileOffset := tileSize + 2

	// todo: there should be a board.draw where the tileImage is passed
	g.tileImage = ebiten.NewImage(tileSize, tileSize)
	screen.Fill(color.Color(color.RGBA{0xCC, 0xCC, 0xCC, 0xFF}))
	g.tileImage.Fill(color.Color(color.RGBA{0x00, 0x00, 0x00, 0xFF}))

	for row := 0; row < g.board.height; row++ {
		for column := 0; column < g.board.width; column++ {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(column*tileOffset), float64(row*tileOffset))
			screen.DrawImage(g.tileImage, options)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game, err := NewGame()
	if err = ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
