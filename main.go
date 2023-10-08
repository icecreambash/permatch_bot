package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"lwjal/master/connector"
	"lwjal/master/query"
	"os"
	"strings"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env not present")
	}
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Т-ЧМ-23-1", "set_group 2"),
	),
)

func main() {

	token, exists := os.LookupEnv("BOT_TOKEN")

	if !exists {
		log.Fatal("TOKEN_BOT has been not found.")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.Text = "Здравствуйте, выберите группу в которой вы обучаетесь"
				msg.ReplyMarkup = numericKeyboard

				connector.Group()
				_, err := bot.Send(msg)
				if err != nil {
					return
				}
			}

			switch update.Message.Text {
			case "Мое расписание":
				go query.Group(bot, update)
			case "Расписание на сегодня":
				go query.GetScheduleByIndex(bot, update)
			}

		}

		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			_, err2 := bot.Request(callback)
			if err2 != nil {
				return
			}

			//var callbackData utils.CallbackEvent
			prepareData := strings.Split(update.CallbackQuery.Data, " ")

			switch prepareData[0] {
			case "get_schedule":
				go query.SchemeT(bot, update, prepareData[1], update.CallbackQuery.Message.Chat.ID)
				break
			case "set_group":
				go query.SetGroup(bot, update)
				break
			}
		}
	}
}
