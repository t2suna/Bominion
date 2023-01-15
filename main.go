package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

// 開発用
var diamond Jewel
var farm Action

func init() {

	// 開発用
	diamond = Jewel{
		Name:         "Diamond",
		Price:        0,
		Value:        1,
		VictoryPoint: 1,
	}
	farm = Action{
		Name:       "Farm",
		Price:      2,
		PlusDraw:   1,
		PlusAction: 1,
		PlusBuy:    1,
		PlusValue:  1,
	}

	//本来はconfig inputする
	var playernum int = 2
	for i := 0; i < playernum; i++ {
		//Playerの情報入力を促す
		Players = append(Players, Player{
			Number:      i + 1,
			Name:        "Player" + strconv.Itoa(i+1),
			Score:       0,
			ActionPoint: 1,
			BuyPoint:    1,
			ValuePoint:  0,
		})
		fmt.Println(Players[i].Name)
		Players[i].Init()
		Players[i].PrintHand()
	}

}

func (g *Game) Update() error {
	//メインループ
	//
	if !EndFlag {

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
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
