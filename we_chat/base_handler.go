package weChat

import (
	"github.com/eatmoreapple/openwechat"
	"github.com/web3ssc/ss-go-common/log"
)

type MsgHandlerInterface interface {
	handle(*openwechat.Message) error
	ReplyText(*openwechat.Message) error
}


const (
	GroupHandler = "group"
)

var handlers map[string]MsgHandlerInterface

func init() {
	handlers = make(map[string]MsgHandlerInterface)
	handlers[GroupHandler] = NewGroupMsgHandler()
}

func Handler(msg *openwechat.Message) {
	err := handlers[GroupHandler].handle(msg)
	if err != nil {
		log.Errorf(err, "Handle error")
		return
	}
}