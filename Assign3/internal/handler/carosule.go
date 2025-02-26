package handler

import (
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func HandlerCarousel(bot *linebot.Client, replyToken string) {
	log.Println("HandlerCarousel called")

	carouselTemplate := linebot.NewTemplateMessage(
		"ตัวเลือกสินค้า",
		linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				"https://blog.sahathaishoponline.com/wp-content/uploads/2020/07/Siam77-50-1198x800.jpg",
				"ต้มจืดเงาะ",
				"เงาะยัดไส้หมูสับ ลดขยะอาหาร ดัดแปลงทำให้บริโภคง่ายขึ้น",
				linebot.NewMessageAction("เลือกสินค้า 1", "เลือกสินค้า 1"),
			).WithDefaultAction(
				linebot.NewURIAction("ดูเพิ่มเติม", "https://www.knorr.com/th/r/%25E0%25B8%2595%25E0%25B9%2589%25E0%25B8%25A1%25E0%25B8%2588%25E0%25B8%25B7%25E0%25B8%2594%25E0%25B9%2580%25E0%25B8%2587%25E0%25B8%25B2%25E0%25B8%25B0%25E0%25B8%25AA%25E0%25B8%25AD%25E0%25B8%2594%25E0%25B9%2584%25E0%25B8%25AA%25E0%25B9%2589%25E0%25B8%25AB%25E0%25B8%25A1%25E0%25B8%25B9%25E0%25B8%25AA%25E0%25B8%25B1%25E0%25B8%259A.html/117355"),
			),

			linebot.NewCarouselColumn(
				"https://www.nipponichitaiyaki.com/img/product_1.jpg",
				"Taiyaki ไส้ถั่วแดง",
				"ขนมจากญี่ปุ่น รูปร่างปลา แป้งหอมอร่อย ไส้หวานละมุน",
				linebot.NewMessageAction("เลือกสินค้า 2", "เลือกสินค้า 2"),
			).WithDefaultAction(
				linebot.NewURIAction("ดูเพิ่มเติม", "https://chillchilljapan.com/dictionary/taiyaki-and-imagawayaki/"),
			),
		),
	)

	_, err := bot.ReplyMessage(replyToken, carouselTemplate).Do()
	if err != nil {
		log.Println("Error sending carousel:", err)
	}
}
