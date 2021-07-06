package services

import (
	"quasar-fire-app/interfaces"
	"quasar-fire-app/util"
	"strings"
)

type messagesService struct{}

func NewMessagesService() interfaces.IMessagesService {
	return &messagesService{}
}

func (s *messagesService) GetMessage(messages ...[]string) (msg string) {
	validate := true
	if len(messages) > 1{
		length := len(messages[0])
		for index ,_ := range messages {
			validate = (len(messages[index]) == length) && validate
		}
	}
	if validate && len(messages) > 1 {
		tmp := make([]string, len(messages[0]))
		for i, _ := range messages {
			tmp = util.GetPart(tmp, messages[i])
		}
		return strings.Join(tmp, " ")
	}
	return ""
}

