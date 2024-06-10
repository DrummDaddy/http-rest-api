package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/drummdaddy/http-rest-api/cmd/apiserver/internal/app/apiserver"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	// Загрузка конфигурации для вашего сервера
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Запуск вашего сервера
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

	// Настройка и запуск Telegram бота
	bot, err := tgbotapi.NewBotAPI("7252926007:AAHCU4sIhFttu0zM7z9XDmUEratdFcd2fzI")
	if err != nil {
		log.Panic(err)
	}

	// Устанавливаем Webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://7252926007:AAHCU4sIhFttu0zM7z9XDmUEratdFcd2fzI/webhook"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/webhook")
	go http.ListenAndServe("0.0.0.0:8080", nil)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Это ответ на ваше сообщение.")
		bot.Send(msg)
	}
}
