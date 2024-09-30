package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Game struct {
	boardImage *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	g.boardImage = ebiten.NewImage(20, 20)
	screen.Fill(color.Color(color.RGBA{0xCC, 0xCC, 0xCC, 0xFF}))
	g.boardImage.Fill(color.Color(color.RGBA{0x00, 0x00, 0x00, 0xFF}))

	screen.DrawImage(g.boardImage, nil)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(22), float64(0))
	screen.DrawImage(g.boardImage, options)

	options = &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(0), float64(22))
	screen.DrawImage(g.boardImage, options)

	options = &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(22), float64(22))
	screen.DrawImage(g.boardImage, options)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
