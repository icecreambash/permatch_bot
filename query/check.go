package query

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
)

func GetScheduleByIndex(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	weekday := time.Now().Weekday()

	index := int(weekday)

	master := 0

	if index == 0 || index == 6 {
		master = 1
	} else {
		master = index
	}

	SchemeT(bot, update, strconv.Itoa(master), update.Message.Chat.ID)
}
