package main

type Card interface {
	TellMyName() string
	Activate(*Player)
}

type Jewel struct {
	Name         string
	Price        int
	Value        int
	VictoryPoint int
}

func (j Jewel) TellMyName() string {
	return j.Name
}

func (j Jewel) Activate(p *Player) {

}

type Action struct {
	Name       string
	Price      int
	DrawPlus   int
	ActionPlus int
	BuyPlus    int
	ValuePlus  int
	Special    int
}

func (a Action) TellMyName() string {
	return a.Name
}

func (a Action) ShowText() {
	//効果を見る
}

func (a Action) Activate(p *Player) {
	p.ActionPoint -= 1
	if a.ActionPlus > 0 {
		p.ActionPoint += a.ActionPlus
	}
	if a.BuyPlus > 0 {
		p.BuyPoint += a.BuyPlus
	}
	if a.ValuePlus > 0 {
		p.ValuePoint += a.ValuePlus
	}
	if a.DrawPlus > 0 {
		for i := 0; i < a.DrawPlus; i++ {
			p.DrawHand()
		}
	}
	SpecialActivate(a.Special)

}

/*
	ドロー購入プラス金貨以外の効果が発現する際に使用
*/
func SpecialActivate(sp int) {
	switch sp {
	case 0:
		//なんもしない
	case 1:
		//なんか効果
	}
}
