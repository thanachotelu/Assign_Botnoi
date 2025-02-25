package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func HandlerText(bot *linebot.Client, replyToken string, userID string, message string) {
	var messages []linebot.SendingMessage
	message = strings.ToLower(message)

	switch message {
	case "yo":
		messages = append(messages,
			linebot.NewTextMessage("Yoo! ♥ "),
			linebot.NewTextMessage("เป็นยังไงบ้างช่วงนี้"),
		)
	case "ทำอะไรอยู่":
		messages = append(messages,
			linebot.NewTextMessage("กำลังเล่นเกมอยู่ครับ "),
			linebot.NewTextMessage("มีอะไรให้ช่วยไหม"),
		)
	case "สวัสดี":
		profile, err := bot.GetProfile(userID).Do()
		if err != nil {
			log.Println("Error user profile: ", err)
		}
		messages = append(messages, linebot.NewTextMessage(fmt.Sprintf("สวัสดีนะครับ %s!", profile.DisplayName)))
	case "hi":
		messages = append(messages, linebot.NewTextMessage("Hi!"))
	default:
		messages = append(messages, linebot.NewTextMessage("ขอโทษด้วย ผมไม่เข้าใจ"))
	}

	_, err := bot.ReplyMessage(replyToken, messages...).Do()
	if err != nil {
		log.Println("Error sending messages: ", err)
	}
}
