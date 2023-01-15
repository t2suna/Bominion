package main

type Card interface {
	TellMyName() string
	Activate()
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

func (j Jewel) Activate() {

}

type Action struct {
	Name       string
	Price      int
	PlusDraw   int
	PlusAction int
	PlusBuy    int
	PlusValue  int
	Special    int
}

func (a Action) TellMyName() string {
	return a.Name
}

func (a Action) ShowText() {
	//効果を見る
}

func (a Action) Activate() {
	if a.PlusAction > 0 {

	}
	if a.PlusBuy > 0 {

	}
	if a.PlusValue > 0 {

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
