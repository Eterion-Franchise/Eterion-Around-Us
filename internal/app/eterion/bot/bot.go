package bot

import (
	"eterion_around_us/internal/app/eterion/errors"
	"eterion_around_us/internal/app/eterion/types"
	"log"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func Init() {
	bot, err := telego.NewBot(os.Getenv("TOKEN"), telego.WithDefaultDebugLogger())
	if err != nil {
		panic("Unable to start bot:" + err.Error())
	}
	defer bot.StopLongPolling()

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, err := bot.SendMessage(setMessageParams(
			update.Message.Chat.ChatID(),
			"<i>Вы в архиве Этериона. Какое знание вы хотите открыть сегодня?</i>",
			KeyboardMainMenu,
		))
		if err != nil {
			log.Print(err)
		}
	}, th.CommandEqual("start"))

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		switch message.Text {
		case CampaignsButton.Text:
			_, err := bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				"<b>📖 Хронология событий 📖</b>\n\n1. Абоба",
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case MapsButton.Text:
			_, err := bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				"<b>🗺️ Карты 🗺️</b>",
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case BattlesButton.Text:
			_, err := bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				"<b>⚔️ Битвы ⚔️</b>",
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case MusicButton.Text:
			_, err := bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				"<b>🎵 Плейлисты 🎵</b>",
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case CharButton.Text:
			//
		}
	})

	bh.Start()
}

func formDataString(dataToGet types.WikiDataType) (string, error) {
	var dataString string

	switch dataToGet {
	case types.CAMPAIGNS:
		// db request
		return dataString, nil
	case types.MAPS:
		// db request
		return dataString, nil
	case types.BATTLES:
		// db request
		return dataString, nil
	case types.MUSIC:
		// db request
		return dataString, nil
	default:
		dataString = "<i>Зыбучие северные ветра проносятся мимо вас, как вы слышите глухой отзвук Пустоты...</i>"
		return dataString, &errors.INVALID_FORM_STRING_TYPE
	}
}

func setMessageParams(chatID telego.ChatID,
	text string,
	replyMarkup telego.ReplyMarkup) *telego.SendMessageParams {
	return &telego.SendMessageParams{
		ChatID:      chatID,
		Text:        text,
		ParseMode:   "html",
		ReplyMarkup: replyMarkup,
	}
}
