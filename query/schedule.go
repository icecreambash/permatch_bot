package query

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var scheduleKb = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Понедельник", "get_schedule 1"),
		tgbotapi.NewInlineKeyboardButtonData("Вторник", "get_schedule 2"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Среда", "get_schedule 3"),
		tgbotapi.NewInlineKeyboardButtonData("Четверг", "get_schedule 4"),
		tgbotapi.NewInlineKeyboardButtonData("Пятница", "get_schedule 5"),
	),
)

func Group(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("Расписание")

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите день недели")

	msg.ReplyMarkup = scheduleKb

	_, err := bot.Send(msg)
	if err != nil {
		return
	}
}
