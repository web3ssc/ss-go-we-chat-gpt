package weChat

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/web3ssc/ss-go-common/log"
	"os"
)

func StartWeChat() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = Handler
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	if err := bot.Login(); err != nil {
		log.Errorf(err, "bot.Login error")
		os.Exit(0)
		return
	}

	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Errorf(err, "bot.GetCurrentUser error")
		os.Exit(0)
		return
	}

	friends, err := self.Friends()
	for i, friend := range friends {
		log.Infof("i:%v, fiend:%v", i, friend)
	}

	groups, err := self.Groups()
	for i, group := range groups {
		log.Infof("i:%v, group:%v", i, group)
	}

	err = bot.Block()
	if err != nil {
		log.Errorf(err, "bot Block error")
		os.Exit(0)
		return
	}

}
