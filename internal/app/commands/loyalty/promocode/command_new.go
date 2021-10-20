package promocode

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/loyalty"
	"log"
)

func (c *LoyaltyPromocodeCommander) New(inputMessage *tgbotapi.Message) {
	log.Printf("[new handler] [%s] text: %s", inputMessage.From.UserName, inputMessage.Text)
	var promocode loyalty.Promocode
	err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &promocode)
	if err != nil {
		reply := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Wrong json, err: %s", err.Error()))
		_, err := c.bot.Send(reply)
		if err != nil {
			log.Printf("[new handler] error sending reply message to chat - %v", err)
		}
		return
	}
	id, _ := c.service.Create(promocode)
	text := fmt.Sprintf("Created promocode with id %d, promocode %s", id, promocode.Promocode)
	reply := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	_, err = c.bot.Send(reply)
	if err != nil {
		log.Printf("[new handler] error sending reply message to chat - %v", err)
	}

}
