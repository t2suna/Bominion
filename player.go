package main

import "fmt"

const APINIT = 1
const BPINIT = 1
const VPINIT = 0

type Player struct {
	Number      int
	Name        string
	Score       int
	Hand        []Card
	MyDeck      DeckStruct
	ActionPoint int
	BuyPoint    int
	ValuePoint  int
}

//	初期化
func (p *Player) Init() {
	for i := 0; i < 2; i++ {
		p.MyDeck.Deck = append(p.MyDeck.Deck, farm)
	}
	for i := 0; i < 8; i++ {
		p.MyDeck.Deck = append(p.MyDeck.Deck, diamond)
	}
	p.MyDeck.Shuffle()
	for i := 0; i < 5; i++ {
		p.DrawHand()
	}
}

func (p *Player) PrintHand() {
	for _, v := range p.Hand {
		fmt.Println("@:" + v.TellMyName())
	}
}

//	クリーンアップ
func (p *Player) CallMeCleanUpPhase() {
	p.ActionPoint = APINIT
	p.BuyPoint = BPINIT
	p.ValuePoint = VPINIT
	p.MyDeck.CleanUp()
}

// 	手札にカードをドローする
func (p *Player) DrawHand() {
	p.Hand = append(p.Hand, p.MyDeck.Deck[len(p.MyDeck.Deck)-1])
	p.MyDeck.Deck = p.MyDeck.Deck[:len(p.MyDeck.Deck)-1]
}

// 手札からカードを捨てる
func (p *Player) DiscardHand(index int) {
	p.MyDeck.DiscardZone = append(p.MyDeck.DiscardZone, p.Hand[index])
	p.Hand = append(p.Hand[:index-1], p.Hand[index:]...)
}

// 手札からカードを発動する
func (p *Player) ActivateHand(index int) {
	p.MyDeck.ActivateZone = append(p.MyDeck.ActivateZone, p.Hand[index])
	p.Hand = append(p.Hand[:index-1], p.Hand[index:]...)
	p.Hand[index].Activate()
}
