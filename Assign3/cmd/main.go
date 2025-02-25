package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"

	"assign3_line_chatbot/internal/handler"
	"assign3_line_chatbot/internal/model"
	"assign3_line_chatbot/internal/tokenconfig"
)

var bot *linebot.Client

func main() {
	cfg := tokenconfig.LoadConfig()

	bot, err := linebot.New(cfg.ChannelSecret, cfg.ChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		events, err := bot.ParseRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		webhookEvent := model.LineWebhookEvent{Events: events}

		for _, event := range webhookEvent.Events {
			log.Println("Processing event:", event.Type)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					text := message.Text
					text = strings.ToLower(text)
					switch text {
					case "button":
						handler.HandlerButton(bot, event.ReplyToken)
						return
					case "quickreply":
						handler.HandlerQuickReply(bot, event.ReplyToken)
						return
					case "carousel":
						handler.HandlerCarousel(bot, event.ReplyToken)
						return
					default:
						handler.HandlerText(bot, event.ReplyToken, event.Source.UserID, text)
					}
					log.Println("User message:", message.Text)
				}
			} else if event.Type == linebot.EventTypePostback {
				postbackData := event.Postback.Data
				handler.HandlerImageResponse(bot, event.ReplyToken, postbackData)
				return
			}
		}
	})

	port := cfg.Port //Load Port from config
	if port == "" {
		port = "8081"
	}

	fmt.Println("Server started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
