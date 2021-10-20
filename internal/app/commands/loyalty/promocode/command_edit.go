package promocode

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/loyalty"
	"log"
	"strconv"
	"strings"
)

func (c *LoyaltyPromocodeCommander) Edit(inputMessage *tgbotapi.Message) {
	log.Printf("[edit handler] [%s] text: %s", inputMessage.From.UserName, inputMessage.Text)

	commandArguments := inputMessage.CommandArguments()
	argumentsList := strings.SplitN(commandArguments, " ", 2)
	var reply tgbotapi.MessageConfig

	if len(argumentsList) != 2 {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Wrong number of arguments, expected 2"))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[edit handler] error sending reply message to chat - %v", err)
		}
		return
	}

	promocodeId, err := strconv.ParseUint(argumentsList[0], 10, 64)
	if err != nil {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("First argument doesn't look like numeric id"))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[edit handler] error sending reply message to chat - %v", err)
		}
		return
	}

	var promocode loyalty.Promocode
	log.Printf("got message: %s", inputMessage.Text)
	err = json.Unmarshal([]byte(argumentsList[1]), &promocode)
	if err != nil {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Wrong json. Error: %s", err.Error()))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[edit handler] error sending reply message to chat - %v", err)
		}
		return
	}

	err = c.service.Update(promocodeId, promocode)
	if err != nil {
		reply = tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Unable to update promocode. Error: %s", err.Error()))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[edit handler] error sending reply message to chat - %v", err)
		}
		return
	}
	text := fmt.Sprintf("Edited promocode with id %d", promocodeId)
	reply = tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	_, err = c.bot.Send(reply)
	if err != nil {
		log.Printf("[edit handler] error sending reply message to chat - %v", err)
	}

}
