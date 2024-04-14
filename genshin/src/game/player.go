package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModCard   *ModCard
}

// 玩家测试
func NewTestPlayer() *Player {
	//玩家模块初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModCard = new(ModCard)

	//数据初始化
	//player.ModPlayer.Icon = 0
	//
	return player
}

// 对外接口
// Icon
func (self *Player) RecvSetIcon(iconId int) {
	self.ModPlayer.SetIcon(iconId, self)
}
// Card
func (self *Player) RecvSetCard(cardId int) {
	self.ModPlayer.SetCard(cardId, self)
}
// Name
func (self *Player) RecvSetName(name string) {
	self.ModPlayer.SetName(name, self)
}
// Card
func (self *Player) RecvSetSign(sign string) {
	self.ModPlayer.SetSign(sign, self)
}
