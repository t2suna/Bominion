package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	WhosFirst(Players)

}

func (g *Game) Update() error {

	//メインループ
	if EndFlag {
		//終了処理
	}

	switch Phase {
	//アクションフェーズ
	case ActionPhase:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(Players[WhosTurn].Hand) > 0 {
			Players[WhosTurn].ActivateHand(Players[WhosTurn].HandIndex)
			//Debug
			fmt.Println(Players[WhosTurn].Name)
			Players[WhosTurn].PrintHand()
			Players[WhosTurn].HandIndex = 0
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {

			fmt.Println(Players[WhosTurn].Name)

			if Players[WhosTurn].HandIndex < len(Players[WhosTurn].Hand)-1 {
				Players[WhosTurn].HandIndex++

			} else {
				Players[WhosTurn].HandIndex = 0
			}

			for i, v := range Players[WhosTurn].Hand {
				if i == Players[WhosTurn].HandIndex {
					fmt.Println("->" + Players[WhosTurn].Hand[Players[WhosTurn].HandIndex].TellMyName())
				} else {
					fmt.Println("@:" + v.TellMyName())
				}
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {

			fmt.Println(Players[WhosTurn].Name)

			if Players[WhosTurn].HandIndex > 0 {
				Players[WhosTurn].HandIndex--

			} else {
				Players[WhosTurn].HandIndex = len(Players[WhosTurn].Hand) - 1
			}

			for i, v := range Players[WhosTurn].Hand {
				if i == Players[WhosTurn].HandIndex {
					fmt.Println("->" + Players[WhosTurn].Hand[Players[WhosTurn].HandIndex].TellMyName())
				} else {
					fmt.Println("@:" + v.TellMyName())
				}
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
