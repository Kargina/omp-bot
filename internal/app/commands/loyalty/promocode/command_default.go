package promocode

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *LoyaltyPromocodeCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[default handler] [%s] text: %s", inputMessage.From.UserName, inputMessage.Text)

	reply := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("you wrote: %s", inputMessage.Text))

	_, err := c.bot.Send(reply)
	if err != nil {
		log.Printf("[default handler] error sending reply message to chat - %v", err)
	}
}
