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
var imgBook []*ebiten.Image
var hand []int

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
	imgBook = append(imgBook, img)
	imgBook = append(imgBook, img2)
	selectImg, _, err = ebitenutil.NewImageFromFile("light.png")
	if err != nil {
		log.Fatal(err)
	}
	// テスト用手札
	hand = append(hand, 0, 1, 1, 1, 0)

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

type countLogic struct {
	l int
	r int
}

var cl countLogic

func (g *Game) MoreDraw(x, y float64, img *ebiten.Image, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	op.GeoM.Scale(0.5, 0.5)

	screen.DrawImage(img, op)
}

const countLimit int = 70

func (g *Game) Draw(screen *ebiten.Image) {

	for i, j := range hand {
		g.MoreDraw(750*float64(i), 500, imgBook[j], screen)
	}

	y = 500
	op := &ebiten.DrawImageOptions{}
	if ebiten.IsKeyPressed(ebiten.KeyE) && cl.l > countLimit && x < 750*4 {
		x += 750
		cl.l = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		cl.l++
	} else if ebiten.IsKeyPressed(ebiten.KeyR) && cl.r > countLimit && x > 0 {
		x -= 750
		cl.r = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
		cl.r++
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
