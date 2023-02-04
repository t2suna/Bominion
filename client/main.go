package client

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

// 開発用

var playernum int
var img *ebiten.Image
var img2 *ebiten.Image

func init() {

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
var countlogic int

func (g *Game) Draw(screen *ebiten.Image) {
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
	//screen.DrawImage(img, op)

	screen.DrawImage(img2, op)
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
