package promocode

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *LoyaltyPromocodeCommander) Delete(inputMessage *tgbotapi.Message) {
	log.Printf("[delete handler] [%s] text: %s", inputMessage.From.UserName, inputMessage.Text)
	commandArguments := inputMessage.CommandArguments()
	argumentsList := strings.Split(commandArguments, " ")
	var reply tgbotapi.MessageConfig

	if len(argumentsList) != 1 {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Wrong number of arguments, expected 1 - promocode id"))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[delete handler] error sending reply message to chat - %v", err)
		}
		return
	}
	promocodeId, err := strconv.ParseUint(argumentsList[0], 10, 64)
	if err != nil {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Argument doesn't look like numeric id"))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[delete handler] error sending reply message to chat - %v", err)
		}
		return
	}

	deleted, err := c.service.Remove(promocodeId)

	if deleted {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Promocode with id %d deleted", promocodeId))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[delete handler] error sending reply message to chat - %v", err)
		}
		return
	} else {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Error: %s", err.Error()))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[delete handler] error sending reply message to chat - %v", err)
		}
		return
	}
}
