package main

import (
	"fmt"
	"otaku/src/csvs"
	"otaku/src/game"
	"time"
)

func main() {
	// 初始化,加载配置
	csvs.CheckLoadCsv()
	//server start
	fmt.Println("数据测试----------start")
	go game.GetManageBanWord().Run()
	tickerIn := time.NewTicker(time.Second * 3)
	tickerOut := time.NewTicker(time.Second * 5)
	player := game.NewTestPlayer()
	player.RecvSetIcon(1)

	for {
		select {
		case <-tickerIn.C:
			player.RecvSetIcon(int(time.Now().Unix()))
		case <-tickerOut.C:
			player.RecvSetName("shon")
		}
	}

	/*
	   	player:= game.NewTestPlayer()
	   	gm := game.NewTestPlayer()
	   	//设置图标测试
	   	fmt.Println("-----玩家1------")
	   	//fmt.Println(player)
	   	player.RecvSetName("好人")
	   	player.RecvSetName("丘外挂")
	   	 player.RecvSetIcon(0)
	   	player.RecvSetIcon(1) //胡桃
	   	player.RecvSetIcon(2) //胡桃
	   	player.RecvSetCard(11)
	   	player.RecvSetCard(22)
	   	player.RecvSetCard(33)
	   	fmt.Println("-------GM------")
	   //	fmt.Println(gm)
	   	gm.RecvSetIcon(2) // 温迪
	   	gm.RecvSetIcon(3) //钟离
	*/

}
