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
var playernum int

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
		DrawPlus:   1,
		ActionPlus: 1,
		BuyPlus:    1,
		ValuePlus:  1,
	}

	//本来はconfig inputする
	playernum = 2
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
	if EndFlag {
		//終了処理
	}

	switch Phase {
	//アクションフェーズ
	case ActionPhase:
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			Players[WhosTurn-1].ActivateHand(Players[WhosTurn-1].HandIndex)
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			if Players[WhosTurn-1].HandIndex < len(Players[WhosTurn-1].Hand) {
				Players[WhosTurn-1].HandIndex++
			}
		}
	case BuyPhase:
	case CleanUpPhase:
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
