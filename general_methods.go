package telebot

import (
	"fmt"
	"errors"
	"encoding/json"
)

/*****************************************************************************/

func (bot *Bot) GetMe() (user *User, err error) {
	url := bot.fullUrl + "/getMe"
	body, err := bot.getRequestWithBody(url)
	if err != nil {
		err = errors.New("Bot.GetMe(): " + err.Error())
		return nil, err
	}
	type	getMeJSON	struct {
		Ok 				bool
		Result			User
	}
	var respJSON getMeJSON
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		err = errors.New("Bot.GetMe(): " + err.Error())
		return nil, err
	}
	return &respJSON.Result, nil
}

/*****************************************************************************/

// TODO: make sendMessage more elegant

type	MsgParams		struct {
	Chat_id						int
	Text						string
	Parse_mode					string
	Disable_web_page_preview	bool
	Disable_notification		bool
	Reply_to_message_id			int
	Reply_markup				interface{}
}

func completeSendMessageUrl(p *MsgParams) (url string) {
	url = url + fmt.Sprintf("?chat_id=%d&text=%s", p.Chat_id, p.Text)
	url = url + fmt.Sprintf("&parse_mode=%s", p.Parse_mode)
	url = url + fmt.Sprintf("&disable_web_page_preview=%t", p.Disable_web_page_preview)
	url = url + fmt.Sprintf("&disable_notification=%t", p.Disable_notification)
	url = url + fmt.Sprintf("&reply_to_message_id=%d", p.Reply_to_message_id)
	if (p.Reply_markup != nil) {
		url = url + fmt.Sprintf("&reply_markup=%v", p.Reply_markup)
	}
	fmt.Printf("%s\n",p.Reply_markup)
	return url
}

func (bot *Bot) SendMessage(p MsgParams) (message *Message, err error) {
	url := bot.fullUrl + "/sendMessage" + completeSendMessageUrl(&p)
	body, err := bot.getRequestWithBody(url)
	if err != nil {
		err = errors.New("Bot.SendMessage(): " + err.Error())
		return nil, err
	}
	type	sendMsgJSON	struct {
		Ok 				bool
		Result			Message
	}
	var respJSON sendMsgJSON
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		err = errors.New("Bot.SendMessage(): " + err.Error())
		return nil, err
	}
	return &respJSON.Result, nil
}

/*****************************************************************************/

