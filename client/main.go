package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

// 開発用

var playernum int
var img *ebiten.Image
var img2 *ebiten.Image
var selectImg *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("card.png")
	if err != nil {
		log.Fatal(err)
	}
	img2, _, err = ebitenutil.NewImageFromFile("card2.png")
	if err != nil {
		log.Fatal(err)
	}
	selectImg, _, err = ebitenutil.NewImageFromFile("light.png")
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update() error {
	//メインループ
	if EndFlag {
		//終了処理
	} else {
	}

	switch Phase {
	//アクションフェーズ
	case ActionPhase:

	case BuyPhase:

	case CleanUpPhase:

	}
	return nil
}

var x, y float64
var x2, y2 float64
var countlogic int

func (g *Game) Draw(screen *ebiten.Image) {

	x2 = 0
	y2 = 500

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(x2, y2)
	op2.GeoM.Scale(0.5, 0.5)

	screen.DrawImage(img2, op2)

	y = 500
	op := &ebiten.DrawImageOptions{}
	if ebiten.IsKeyPressed(ebiten.KeyE) && countlogic > 100 && x < 750*4 {
		x += 750
		countlogic = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		countlogic++
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) && countlogic > 100 && x > 0 {
		x -= 750
		countlogic = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
		countlogic++
	}

	op.GeoM.Translate(x, y)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(selectImg, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 2000, 1200
}

func main() {

	ebiten.SetWindowSize(1000, 600)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
