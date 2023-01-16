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

	NumSupply = map[Card]int{}
	Supply = append(Supply, diamond)
	Supply = append(Supply, farm)
	NumSupply[diamond] = NumJewelSupply
	NumSupply[farm] = NumActionSupply

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
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			Phase = BuyPhase
			Players[WhosTurn].Pointer = 0
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(Players[WhosTurn].Hand) > 0 {
			Players[WhosTurn].ActivateHand(Players[WhosTurn].Pointer)
			//Debug
			fmt.Println(Players[WhosTurn].Name)
			Players[WhosTurn].PrintHand()
			Players[WhosTurn].Pointer = 0
			if Players[WhosTurn].ActionPoint == 0 {
				Phase = BuyPhase
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {

			fmt.Println(Players[WhosTurn].Name)

			if Players[WhosTurn].Pointer < len(Players[WhosTurn].Hand)-1 {
				Players[WhosTurn].Pointer++

			} else {
				Players[WhosTurn].Pointer = 0
			}

			for i, v := range Players[WhosTurn].Hand {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {

			fmt.Println(Players[WhosTurn].Name)

			if Players[WhosTurn].Pointer > 0 {
				Players[WhosTurn].Pointer--

			} else {
				Players[WhosTurn].Pointer = len(Players[WhosTurn].Hand) - 1
			}

			for i, v := range Players[WhosTurn].Hand {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
		}
	case BuyPhase:

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			Phase = CleanUpPhase
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(Players[WhosTurn].Hand) > 0 {
			Players[WhosTurn].BuyCard(Supply[Players[WhosTurn].Pointer])
			Players[WhosTurn].Pointer = 0
			if Players[WhosTurn].BuyPoint == 0 {
				Phase = CleanUpPhase
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			fmt.Println("Supply")

			if Players[WhosTurn].Pointer < len(Supply)-1 {
				Players[WhosTurn].Pointer++

			} else {
				Players[WhosTurn].Pointer = 0
			}
			for i, v := range Supply {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			fmt.Println("Supply")

			if Players[WhosTurn].Pointer > 0 {
				Players[WhosTurn].Pointer--

			} else {
				Players[WhosTurn].Pointer = len(Supply) - 1
			}

			for i, v := range Supply {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
		}
	case CleanUpPhase:
		Players[WhosTurn].CallMeCleanUpPhase()
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
