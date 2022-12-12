package weChat

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/web3ssc/ss-go-common/log"
	"github.com/web3ssc/ss-go-we-chat-gpt/config"
	"github.com/web3ssc/ss-go-we-chat-gpt/open_ai"
	"strings"
)

type GroupMsgHandler struct {

}

func NewGroupMsgHandler() *GroupMsgHandler {
	return &GroupMsgHandler{}
}

func (handler *GroupMsgHandler) handle(msg *openwechat.Message) error{
	if !msg.IsText() {
		return nil
	}
	return handler.ReplyText(msg)
}

func (handler *GroupMsgHandler) ReplyText(msg *openwechat.Message) error {
	sender, err := msg.Sender()
	group := openwechat.Group{User: sender}
	log.Infof("from: %v, msg: %v", group.NickName, msg.Content)
	if !strings.Contains(msg.Content, config.Nickname) {
		return nil
	}
	splitItems := strings.Split(msg.Content, config.Nickname)
	if len(splitItems) < 2 {
		return nil
	}
	requestText := strings.TrimSpace(splitItems[1])
	reply, err := open_ai.Completions(requestText)
	if err != nil {
		log.Errorf(err, "open_ai.Completions error")
		_, err := msg.ReplyText(fmt.Sprintf("机器人出错啦，请稍后再试"))
		if err != nil {
			return err
		}
		return err
	}

	if reply != nil {
		log.Infof("reply: %v", reply)
		_, err = msg.ReplyText(*reply)
		if err != nil {
			log.Errorf(err, "msg.ReplyText error")
		}
		return err
	}

	return nil
}