package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
}

func NewTestPlayer() *Player {
	//玩家初始化
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)

	//数据初始化
	player.ModPlayer.Icon = 0
	//
	return player
}
