package query

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var groupKb = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Мое расписание"),
		tgbotapi.NewKeyboardButton("Расписание на сегодня"),
	),
)

func SetGroup(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вы успешно прикрепили группу")

	msg.ReplyMarkup = groupKb

	_, err := bot.Send(msg)
	if err != nil {
		return
	}

}
