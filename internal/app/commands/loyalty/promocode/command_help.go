package promocode

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LoyaltyPromocodeCommander) Help(inputMessage *tgbotapi.Message) {
	log.Printf("[help handler] [%s] text: %s", inputMessage.From.UserName, inputMessage.Text)
	reply := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__loyalty__promocode - help\n"+
			"/list__loyalty__promocode - list of promocodes\n"+
			"/get__loyalty__promocode - information about promocode\n"+
			"/delete__loyalty__promocode - delete promocode\n"+
			"/new__loyalty__promocode - create new promocode\n"+
			"/edit__loyalty__promocode - update information about promocode\n",
	)

	_, err := c.bot.Send(reply)
	if err != nil {
		log.Printf("[help handler] error sending reply message to chat - %v", err)
	}
}
