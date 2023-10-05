package main

import (
	"github.com/Rustamchuk/botRustam/internal/app/commands"
	"github.com/Rustamchuk/botRustam/internal/speaking/phrases"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

const TOKEN = "5970472266:AAGDeQx-uS-hXqgYKyu2pK8eo9qBpdg5C_8"

func main() {
	//godotenv.Load()
	//
	//token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	tangue := phrases.NewService()

	commander := commands.NewCommander(bot, tangue)

	for update := range updates {
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		commander.ChooseMove(&update)
	}
}
