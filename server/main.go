package server

import (
	"fmt"
	_ "image/png"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	sys "github.com/bominion/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

// 開発用
var playernum int
var img *ebiten.Image
var img2 *ebiten.Image

func init() {

	// 開発用
	sys.Diamond = sys.Jewel{
		Name:         "Diamond",
		Price:        0,
		Value:        1,
		VictoryPoint: 1,
	}
	sys.Farm = sys.Action{
		Name:       "Farm",
		Price:      2,
		DrawPlus:   1,
		ActionPlus: 1,
		BuyPlus:    1,
		ValuePlus:  1,
	}

	sys.NumSupply = map[sys.Card]int{}
	sys.Supply = append(sys.Supply, sys.Diamond)
	sys.Supply = append(sys.Supply, sys.Farm)
	sys.NumSupply[sys.Diamond] = sys.NumJewelSupply
	sys.NumSupply[sys.Farm] = sys.NumActionSupply

	//本来はconfig inputする
	playernum = 2
	for i := 0; i < playernum; i++ {
		//Playerの情報入力を促す
		sys.Players = append(sys.Players, sys.Player{
			Number:      i + 1,
			Name:        "Player" + strconv.Itoa(i+1),
			Score:       0,
			ActionPoint: 1,
			BuyPoint:    1,
			ValuePoint:  0,
		})
		sys.Players[i].Init()
		sys.Players[i].PrintHand(0)
	}

	sys.WhosFirst(sys.Players)

}

func (g *Game) Update() error {
	//fmt.Println(ebiten.ActualFPS())
	//fmt.Println(ebiten.ActualTPS())
	//メインループ

	switch sys.Phase {
	//アクションフェーズ
	case sys.ActionPhase:
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			sys.Phase = sys.BuyPhase
			sys.Players[sys.WhosTurn].Pointer = 0

			//fmt.Println("Supply")
			/*
				for i, v := range sys.Supply {
					if i == sys.Players[sys.WhosTurn].Pointer {
						fmt.Println("->" + v.TellMyName())
					} else {
						fmt.Println("@" + v.TellMyName())
					}
				}
			*/
			//fmt.Println("__________")
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(sys.Players[sys.WhosTurn].Hand) > 0 {
			sys.Players[sys.WhosTurn].ActivateHand(sys.Players[sys.WhosTurn].Pointer)
			//sys.Players[sys.WhosTurn].PrintHand(0)
			//sys.Players[sys.WhosTurn].Pointer = 0

			if sys.Players[sys.WhosTurn].ActionPoint == 0 {
				sys.Phase = sys.BuyPhase
				/*
					fmt.Println("Supply")
					for i, v := range sys.Supply {
						if i == sys.Players[sys.WhosTurn].Pointer {
							fmt.Println("->" + v.TellMyName())
						} else {
							fmt.Println("@" + v.TellMyName())
						}
					}
				*/
				//fmt.Println("__________")
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			/*
				if sys.Players[sys.WhosTurn].Pointer < len(sys.Players[sys.WhosTurn].Hand)-1 {
					sys.Players[sys.WhosTurn].Pointer++

				} else {
					sys.Players[sys.WhosTurn].Pointer = 0
				}

				sys.Players[sys.WhosTurn].PrintHand(sys.Players[sys.WhosTurn].Pointer)
			*/

		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			/*

				if sys.Players[sys.WhosTurn].Pointer > 0 {
					sys.Players[sys.WhosTurn].Pointer--

				} else {
					sys.Players[sys.WhosTurn].Pointer = len(sys.Players[sys.WhosTurn].Hand) - 1
				}

				sys.Players[sys.WhosTurn].PrintHand(sys.Players[sys.WhosTurn].Pointer)
			*/
		}
	case sys.BuyPhase:

		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			sys.Phase = sys.CleanUpPhase
			//sys.Players[sys.WhosTurn].Pointer = 0
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && len(sys.Players[sys.WhosTurn].Hand) > 0 {
			sys.Players[sys.WhosTurn].BuyCard(sys.Supply[sys.Players[sys.WhosTurn].Pointer])
			//sys.Players[sys.WhosTurn].Pointer = 0
			if sys.Players[sys.WhosTurn].BuyPoint == 0 {
				sys.Phase = sys.CleanUpPhase
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
			/*
				if sys.Players[sys.WhosTurn].Pointer < len(sys.Supply)-1 {
					sys.Players[sys.WhosTurn].Pointer++

				} else {
					sys.Players[sys.WhosTurn].Pointer = 0
				}
				+*/
			/*
				fmt.Println("Supply")
				for i, v := range sys.Supply {
					if i == sys.Players[sys.WhosTurn].Pointer {
						fmt.Println("->" + v.TellMyName())
					} else {
						fmt.Println("@" + v.TellMyName())
					}
				}
				fmt.Println("__________")
			*/
		} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
			/*

				if sys.Players[sys.WhosTurn].Pointer > 0 {
					sys.Players[sys.WhosTurn].Pointer--

				} else {
					sys.Players[sys.WhosTurn].Pointer = len(sys.Supply) - 1
				}
			*/
			/*
				fmt.Println("Supply")
				for i, v := range sys.Supply {
					if i == sys.Players[sys.WhosTurn].Pointer {
						fmt.Println("->" + v.TellMyName())
					} else {
						fmt.Println("@" + v.TellMyName())
					}
				}
				fmt.Println("__________")
			*/
		}
	case sys.CleanUpPhase:
		sys.Players[sys.WhosTurn].CallMeCleanUpPhase()
		if sys.WhosTurn < playernum-1 {
			sys.WhosTurn += 1
		} else {
			sys.WhosTurn = 0
		}
		//sys.Players[sys.WhosTurn].PrintHand(0)
		sys.Phase = sys.ActionPhase
	}

	return nil
}

var x, y float64
var countlogic int

func main() {
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	/*Ref from
	https://zenn.dev/hsaki/books/golang-grpc-starting/viewer/server
	*/

	server, _ := NewGRPCServer()
	go func() {
		log.Printf("start gRPC server port: %v", port)
		server.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()

}
