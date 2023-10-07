package query

import (
	"bytes"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"html/template"
	"lwjal/master/utils"
	"os"
)

const htmlTemplate = `
	<b>{{.FullCode}}</b>
`

func SchemeT(bot *tgbotapi.BotAPI, update tgbotapi.Update, id string) {

	for _, element := range utils.TimeSchedule {
		if element.Code == id {
			webpage, _ := template.New("template").Parse(htmlTemplate)

			var _ = webpage.Execute(os.Stdout, element)

			/*
				Тут собираем расписание
			*/
			var rows [][]tgbotapi.InlineKeyboardButton

			for _, item := range element.Items {
				fmt.Println(item)

				item := [][]tgbotapi.InlineKeyboardButton{
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData(item.Name, "test"),
					),
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData(item.Time+" || "+item.Class, "test"),
					),
				}

				rows = append(rows, item...)

			}

			var tpl bytes.Buffer
			webpage.Execute(&tpl, element)

			result := tpl.String()

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, result)

			msg.ParseMode = tgbotapi.ModeHTML

			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(rows...)

			_, err := bot.Send(msg)
			if err != nil {
				return
			}

		}
	}
}
