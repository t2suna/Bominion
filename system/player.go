package system

import "fmt"

const APINIT = 1
const BPINIT = 1
const VPINIT = 0

type Player struct {
	Number      int
	Name        string
	Score       int
	Hand        []Card
	Pointer     int
	MyDeck      DeckStruct
	ActionPoint int
	BuyPoint    int
	ValuePoint  int
}

//	初期化
func (p *Player) Init() {
	for i := 0; i < 2; i++ {
		p.MyDeck.Deck = append(p.MyDeck.Deck, Farm)
	}
	for i := 0; i < 8; i++ {
		p.MyDeck.Deck = append(p.MyDeck.Deck, Diamond)
	}
	p.MyDeck.Shuffle(p.Number)
	for i := 0; i < 5; i++ {
		p.DrawHand()
	}
}

func (p *Player) PrintHand(pointer int) {
	fmt.Println(p.Name)
	for i, v := range p.Hand {
		if i == pointer {
			fmt.Println("->" + v.TellMyName())
		} else {
			fmt.Println("@" + v.TellMyName())
		}
	}
	fmt.Println("__________")
}

//	クリーンアップ
func (p *Player) CallMeCleanUpPhase() {
	p.DiscardAllHand()
	p.ActionPoint = APINIT
	p.BuyPoint = BPINIT
	p.ValuePoint = VPINIT
	p.MyDeck.CleanActivateZone()
	for i := 0; i < 5; i++ {
		p.DrawHand()
	}
}

// 	カードを購入する
func (p *Player) BuyCard(card Card) {
	p.BuyPoint--
	if card.TellMyPrice() <= p.ValuePoint {
		p.ValuePoint -= card.TellMyPrice()
		p.MyDeck.DiscardZone = append(p.MyDeck.DiscardZone, card)
		fmt.Println("You bought " + card.TellMyName())
	}
}

// 	手札にカードを一枚ドローする
func (p *Player) DrawHand() {
	if len(p.MyDeck.Deck) == 0 {
		p.MyDeck.CleanDiscardZone()
		p.MyDeck.Shuffle(0)
	}
	p.Hand = append(p.Hand, p.MyDeck.Deck[len(p.MyDeck.Deck)-1])
	p.MyDeck.Deck = p.MyDeck.Deck[:len(p.MyDeck.Deck)-1]
}

// 手札からカードをすべて捨てる
func (p *Player) DiscardAllHand() {
	p.MyDeck.DiscardZone = append(p.MyDeck.DiscardZone, p.Hand...)
	p.Hand = []Card{}
}

// 手札からカードを捨てる
func (p *Player) DiscardHand(index int) {
	p.MyDeck.DiscardZone = append(p.MyDeck.DiscardZone, p.Hand[index])
	p.Hand = append(p.Hand[:index-1], p.Hand[index:]...)
}

// 手札からカードを発動する
func (p *Player) ActivateHand(index int) {

	p.MyDeck.ActivateZone = append(p.MyDeck.ActivateZone, p.Hand[index])
	if index != 0 {
		p.Hand = append(p.Hand[:index-1], p.Hand[index:]...)
	} else if len(p.Hand) > 1 {
		p.Hand = p.Hand[1:]
	} else {
		p.Hand = []Card{}

	}
	p.MyDeck.ActivateZone[len(p.MyDeck.ActivateZone)-1].Activate(p)

}
