package handler

import (
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func HandlerButton(bot *linebot.Client, replyToken string) {
	log.Println("HandlerText called with message")
	buttonTemplate := linebot.NewTemplateMessage(
		"กรุณาเลือก",
		linebot.NewButtonsTemplate(
			"https://avatars.githubusercontent.com/u/115211659?v=4",
			"GitHub Profile",
			"กดปุ่มด้านล่างเพื่อเข้าสู่ GitHub Profile",
			linebot.NewURIAction("Profile", "https://github.com/thanachotelu"),
		),
	)

	_, err := bot.ReplyMessage(replyToken, buttonTemplate).Do()
	if err != nil {
		log.Println("Error sending button template:", err)
	}
}
