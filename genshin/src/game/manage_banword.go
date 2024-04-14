package game

import (
	"fmt"
	"otaku/src/csvs"
	"regexp"
	"time"
)

// 单例模式启动公共类
var manageBanWord *ManageBanWord

func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		// 先写死几个固定值
		manageBanWord.BanWordBase = []string{"外挂", "工具"}
		manageBanWord.BanWordExtra = []string{"原神", "x"}
	}
	return manageBanWord
}

type ManageBanWord struct {
	BanWordBase  []string
	BanWordExtra []string
}

func (self *ManageBanWord) IsBanWord(txt string) bool {
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}
	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}
	return false
}

func (self *ManageBanWord) Run() {
	self.BanWordBase = csvs.GetBanWordBase()
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			// fmt.Println(time.Now().Unix())
			if time.Now().Unix()%10 == 0 {
				fmt.Println("更新词库")
			} else {
				// fmt.Println("待机")
			}

		}
	}
}
