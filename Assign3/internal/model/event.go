package model

import "github.com/line/line-bot-sdk-go/v8/linebot"

//ใช้สำหรับเก็บข้อมูลที่ได้รับจาก LINE Webhook
type LineWebhookEvent struct {
	Events []*linebot.Event `json:"events"`
}
