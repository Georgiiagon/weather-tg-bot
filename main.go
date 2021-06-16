package main

import (
	"fmt"
	"log"
	"os"

	"weather-tg-bot/api"
	"weather-tg-bot/helpers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Saint Petersburg"),
		tgbotapi.NewKeyboardButton("Sortavala"),
		tgbotapi.NewKeyboardButton("Penza"),
	),
)

func main() {
	helpers.LoadEnv()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}
	var weather api.Weather

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /start, /stop and /makas."
		case "start":
			msg.ReplyMarkup = numericKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "stop":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		case "makas":
			msg.Text = "Krasava"
		default:
			msg.Text = "I don't know that command"
		}

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			weather = api.GetByCity(update.Message.Text)
			msg.Text = prepareMessage(&weather)
		}

		if msg.Text == "" {
			continue
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func prepareMessage(weather *api.Weather) string {
	description := ""
	coord := ""
	if len(weather.Weather) != 0 {
		description = ". Desc: " + weather.Weather[0].Description
		coord = ". Coord: " + fmt.Sprintf("%f", weather.Coord.Lat) + "," + fmt.Sprintf("%f", weather.Coord.Lon)
	} else {
		return ""
	}

	return "Where: " + weather.Name + ". Temperature: " + fmt.Sprintf("%.0f", weather.Main.Temp) + description + coord
}
