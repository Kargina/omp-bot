package promocode

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LoyaltyPromocodeCommander) List(inputMessage *tgbotapi.Message) {
	log.Printf("[list handler] [%s] text: %s", inputMessage.From.UserName, inputMessage.Text)
	commandArguments := inputMessage.CommandArguments()
	argumentsList := strings.Split(commandArguments, " ")
	var reply tgbotapi.MessageConfig

	if len(argumentsList) != 2 {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Wrong number of arguments, expected 2"))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[list handler] error sending reply message to chat - %v", err)
		}
		return
	}
	cursor, err := strconv.ParseUint(argumentsList[0], 10, 64)
	if err != nil {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("First argument doesn't look like numeric id"))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[list handler] error sending reply message to chat - %v", err)
		}
		return
	}
	limit, err := strconv.ParseUint(argumentsList[1], 10, 64)
	if err != nil {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Second argument doesn't look like numeric id"))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[list handler] error sending reply message to chat - %v", err)
		}
		return
	}

	outputMsgText := "Here all the promocodes: \n\n"
	promocodes, _ := c.service.List(cursor, limit)
	for _, p := range promocodes {
		outputMsgText += p.Promocode
		outputMsgText += "\n"
	}

	reply = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	_, err = c.bot.Send(reply)
	if err != nil {
		log.Printf("[list handler] error sending reply message to chat - %v", err)
	}
}
