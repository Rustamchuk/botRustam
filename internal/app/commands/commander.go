package commands

import (
	"github.com/Rustamchuk/botRustam/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var commands = map[string]func(c *Commander, inputMessage *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) ChooseMove(update *tgbotapi.Update) {
	com, ok := commands[update.Message.Command()]
	if ok {
		com(c, update.Message)
		return
	}
	c.Default(update.Message)
}
