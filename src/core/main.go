package main

import (
	"fmt"
	"otaku/src/game"
)

func main() {
	//net

	//server start
	fmt.Println("数据测试----------start")

	players := game.NewTestPlayer()
	player := game.NewTestPlayer()
	//设置图标测试
	fmt.Println("-----玩家1------")
	fmt.Println(players)
	fmt.Println(players.ModPlayer)
	players.RecvSetIcon(0)
	player.RecvSetIcon(1) //胡桃
	fmt.Println("-------GM------")
	fmt.Println(player)
	fmt.Println(player.ModPlayer)
	player.RecvSetIcon(2) // 温迪
	player.RecvSetIcon(3) //钟离

}
