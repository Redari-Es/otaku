package game

type ModIcon struct {
}

func (self *ModIcon) IsHasIcon(iconId int) bool {
	if iconId > 0 {
		return false
	} else {
		return true
	}
}
