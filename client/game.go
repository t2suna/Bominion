package main

//後でメインに書くソース

// TODO:グローバル変数にしない
var WhosTurn int = 0
var EndFlag bool = false
var NumActionSupply = 5
var NumJewelSupply = 20

//var Players []Player
var Phase int

//var Supply []Card
//var NumSupply map[Card]int

const (
	WaitPhase   = -1
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
