package main

import (
	"fmt"
	_ "image/png"
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
var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("card.png")
	if err != nil {
		log.Fatal(err)
	}

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
		Players[i].Init()
		Players[i].PrintHand(0)
	}

	WhosFirst(Players)

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
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			Phase = BuyPhase
			Players[WhosTurn].Pointer = 0

			fmt.Println("Supply")
			for i, v := range Supply {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
			fmt.Println("__________")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(Players[WhosTurn].Hand) > 0 {
			Players[WhosTurn].ActivateHand(Players[WhosTurn].Pointer)
			Players[WhosTurn].PrintHand(0)
			Players[WhosTurn].Pointer = 0

			fmt.Println("Supply")
			for i, v := range Supply {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
			fmt.Println("__________")
			if Players[WhosTurn].ActionPoint == 0 {
				Phase = BuyPhase
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {

			if Players[WhosTurn].Pointer < len(Players[WhosTurn].Hand)-1 {
				Players[WhosTurn].Pointer++

			} else {
				Players[WhosTurn].Pointer = 0
			}

			Players[WhosTurn].PrintHand(Players[WhosTurn].Pointer)

		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {

			if Players[WhosTurn].Pointer > 0 {
				Players[WhosTurn].Pointer--

			} else {
				Players[WhosTurn].Pointer = len(Players[WhosTurn].Hand) - 1
			}

			Players[WhosTurn].PrintHand(Players[WhosTurn].Pointer)
		}
	case BuyPhase:

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			Phase = CleanUpPhase
			Players[WhosTurn].Pointer = 0
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(Players[WhosTurn].Hand) > 0 {
			Players[WhosTurn].BuyCard(Supply[Players[WhosTurn].Pointer])
			Players[WhosTurn].Pointer = 0
			if Players[WhosTurn].BuyPoint == 0 {
				Phase = CleanUpPhase
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {

			if Players[WhosTurn].Pointer < len(Supply)-1 {
				Players[WhosTurn].Pointer++

			} else {
				Players[WhosTurn].Pointer = 0
			}
			fmt.Println("Supply")
			for i, v := range Supply {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
			fmt.Println("__________")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {

			if Players[WhosTurn].Pointer > 0 {
				Players[WhosTurn].Pointer--

			} else {
				Players[WhosTurn].Pointer = len(Supply) - 1
			}

			fmt.Println("Supply")
			for i, v := range Supply {
				if i == Players[WhosTurn].Pointer {
					fmt.Println("->" + v.TellMyName())
				} else {
					fmt.Println("@" + v.TellMyName())
				}
			}
			fmt.Println("__________")
		}
	case CleanUpPhase:
		Players[WhosTurn].CallMeCleanUpPhase()
		if WhosTurn < playernum-1 {
			WhosTurn += 1
		} else {
			WhosTurn = 0
		}
		Players[WhosTurn].PrintHand(0)
		Phase = ActionPhase
	}

	return nil
}

var x, y float64
var countlogic int

func (g *Game) Draw(screen *ebiten.Image) {
	y = 1000
	op := &ebiten.DrawImageOptions{}
	if ebiten.IsKeyPressed(ebiten.KeyE) && countlogic == 10 {
		x += 500
		countlogic = 0
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		countlogic++
	}
	op.GeoM.Translate(x, y)
	op.GeoM.Scale(0.1, 0.1)
	screen.DrawImage(img, op)
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
