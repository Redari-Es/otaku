package game

import "fmt"

type ModPlayer struct {
	UserId         int
	Icon           int
	Card           int
	Name           string
	Sign           string
	PlayerLevel    int
	PlayerExp      int
	WorldLevel     int
	WorldLevelCool int64
	Birth          int
	ShowTeam       []int
	ShowCard       int
	//看不见的字段
	IsProhibit int
	IsGM       int
}

func (self *Player) RecvSetIcon(iconId int) {
	fmt.Println("seticon")
	if self.ModIcon.IsHasIcon(iconId) {
		fmt.Println("非法设置图标")
		//通知客户端，操作非法
		return
	}
	self.ModPlayer.Icon = iconId
	fmt.Println("当前图标：", self.ModPlayer.Icon)

}
