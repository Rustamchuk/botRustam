package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong arg", args)
		return
	}
	res, err := c.productService.Get(arg - 1)
	if err != nil {
		log.Println(err.Error())
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Product Title: "+res.Title))
	c.bot.Send(msg)
}
