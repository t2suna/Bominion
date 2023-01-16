package main

//後でメインに書くソース

import (
	"math/rand"
	"time"
)

// TODO:グローバル変数にしない
var WhosTurn int = 0
var EndFlag bool = false
var NumActionSupply = 5
var NumJewelSupply = 20
var Players []Player
var Phase int
var Supply []Card
var NumSupply map[Card]int

const (
	ActionPhase = iota
	BuyPhase
	CleanUpPhase
)

/*
0 アクション(初期)
1 購入
2 クリーンアップ
*/
//乱数はサーバーもしくはホストユーザーのアプリケーションで行う必要がある

func WhosFirst(p []Player) {
	rand.Seed(time.Now().UnixNano())
	WhosTurn = rand.Intn(len(p))
}
