package system

import (
	"math/rand"
	"time"
)

type DeckStruct struct {
	Deck         []Card
	ActivateZone []Card
	DiscardZone  []Card
}

//https://qiita.com/YumaInaura/items/bebb348746978f5abab7
func (d *DeckStruct) Shuffle(Seed int) {
	rand.Seed(time.Now().UnixNano() + int64(Seed))
	rand.Shuffle(len(d.Deck), func(i, j int) { d.Deck[i], d.Deck[j] = d.Deck[j], d.Deck[i] })
}

func (d *DeckStruct) CleanDiscardZone() {
	d.Deck = append(d.Deck, d.DiscardZone...)
	d.DiscardZone = []Card{}
}

func (d *DeckStruct) CleanActivateZone() {
	d.Deck = append(d.DiscardZone, d.ActivateZone...)
	d.ActivateZone = []Card{}
}
