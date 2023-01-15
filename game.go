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
var Players []Player

//乱数はサーバーもしくはホストユーザーのアプリケーションで行う必要がある

func WhosFirst(p []Player) {
	rand.Seed(time.Now().UnixNano())
	WhosTurn = rand.Intn(100)
}
