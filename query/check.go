package query

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func GetScheduleByIndex(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	weekday := time.Now().Weekday()

	index := weekday

	master := "0"

	if index == 0 || index == 6 {
		master = "1"
	} else {
		master = string(index)
	}

	SchemeT(bot, update, master, update.Message.Chat.ID)
}
