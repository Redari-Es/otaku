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

// SetIcon
func (self *ModPlayer) SetIcon(iconId int, player *Player) {
	fmt.Println("setIcon")
	if !player.ModIcon.IsHasIcon(iconId) {
		fmt.Println("非法设置图标")
		//通知客户端，操作非法
		return
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("当前图标：", player.ModPlayer.Icon)
}
// SetCard
func (self *ModPlayer) SetCard(cardId int, player *Player) {
	fmt.Println("setIcon")
	if !player.ModCard.IsHasCard(cardId) {
		fmt.Println("非法设置名片")
		//通知客户端，操作非法
		return
	}
	player.ModPlayer.Card = cardId
	fmt.Println("当前名片：", player.ModPlayer.Card)
}
// SetName
func (self *ModPlayer) SetName(name string, player *Player) {
	fmt.Println("setName")
// 违禁词处理
if GetManageBanWord().IsBanWord(name){
	return
}	
	player.ModPlayer.Name = name
	fmt.Println("当前名字：", player.ModPlayer.Name)
}
// SetSign
func (self *ModPlayer) SetSign(sign string, player *Player) {
	fmt.Println("setSign")
if GetManageBanWord().IsBanWord(sign){
	return
}	
	player.ModPlayer.Sign = sign
	fmt.Println("当前签名：", player.ModPlayer.Sign)
}