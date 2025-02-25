package handler

import (
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func HandlerQuickReply(bot *linebot.Client, replyToken string) {
	log.Println("HandlerQuickReply called with message")
	msg := linebot.NewTextMessage("ชอบสถานที่ไหน").WithQuickReplies(
		linebot.NewQuickReplyItems(
			linebot.NewQuickReplyButton("", linebot.NewPostbackAction("ภูเขา", "mountain", "", "ภูเขา", "", "")),
			linebot.NewQuickReplyButton("", linebot.NewPostbackAction("ทะเล", "sea", "", "ทะเล", "", "")),
			linebot.NewQuickReplyButton("", linebot.NewPostbackAction("แสงเหนือ", "aurora", "", "แสงเหนือ", "", "")),
		),
	)

	_, err := bot.ReplyMessage(replyToken, msg).Do()
	if err != nil {
		log.Println("Error sending quick reply:", err)
	}
}

func HandlerImageResponse(bot *linebot.Client, replyToken string, data string) {
	imageURL := ""

	switch data {
	case "mountain":
		imageURL = "https://images.pexels.com/photos/417173/pexels-photo-417173.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
	case "sea":
		imageURL = "https://images.pexels.com/photos/994605/pexels-photo-994605.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
	case "aurora":
		imageURL = "https://images.pexels.com/photos/360912/pexels-photo-360912.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
	default:
		_, _ = bot.ReplyMessage(replyToken, linebot.NewTextMessage("ไม่พบรูปภาพ")).Do()
	}

	if imageURL != "" {
		msg := linebot.NewImageMessage(imageURL, imageURL)
		_, err := bot.ReplyMessage(replyToken, msg).Do()
		if err != nil {
			log.Println("Error sending image:", err)
		}
	} else {
		_, _ = bot.ReplyMessage(replyToken, linebot.NewTextMessage("ไม่พบรูปภาพ")).Do()
	}
}
