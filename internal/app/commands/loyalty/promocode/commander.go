package promocode

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"

	service "github.com/ozonmp/omp-bot/internal/service/loyalty/promocode"
)

type PromocodeCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type LoyaltyPromocodeCommander struct {
	bot     *tgbotapi.BotAPI
	service service.PromocodeService
}

func NewLoyaltyPromocodeCommander(bot *tgbotapi.BotAPI) *LoyaltyPromocodeCommander {
	promocodeService := service.NewDummyPromocodeService()

	return &LoyaltyPromocodeCommander{
		bot:     bot,
		service: promocodeService,
	}
}

func (c *LoyaltyPromocodeCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LoyaltyPromocodeCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *LoyaltyPromocodeCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "new":
		c.New(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
