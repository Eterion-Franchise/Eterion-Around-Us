package bot

import (
	"eterion_around_us/internal/app/eterion/errors"
	"eterion_around_us/internal/app/eterion/types"
	"fmt"
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
			dataString, err := formDataString(types.CAMPAIGNS)
			if err != nil {
				log.Println(err)
			}
			_, err = bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				dataString,
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case MapsButton.Text:
			dataString, err := formDataString(types.MAPS)
			if err != nil {
				log.Println(err)
			}
			_, err = bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				dataString,
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case BattlesButton.Text:
			dataString, err := formDataString(types.BATTLES)
			if err != nil {
				log.Println(err)
			}
			_, err = bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				dataString,
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case MusicButton.Text:
			dataString, err := formDataString(types.MUSIC)
			if err != nil {
				log.Println(err)
			}
			_, err = bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				dataString,
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		case CharButton.Text:
			dataString, err := formDataString(types.CHARACTER)
			if err != nil {
				log.Println(err)
			}
			_, err = bot.SendMessage(setMessageParams(
				message.Chat.ChatID(),
				dataString,
				KeyboardMainMenu,
			))
			if err != nil {
				log.Print(err)
			}
		}
	})

	bh.Start()
}

func formDataString(dataToGet types.WikiDataType) (string, error) {
	var dataString string

	switch dataToGet {
	case types.CAMPAIGNS:
		// db request
		dataString = fmt.Sprintf("📖 <b>Хронология событий</b> 📖\n\n")
		return dataString, nil
	case types.MAPS:
		// db request
		dataString = fmt.Sprintf("🗺️ <b>Карты</b> 🗺️\n\n")
		return dataString, nil
	case types.BATTLES:
		// db request
		dataString = fmt.Sprintf("⚔️ <b>Битвы</b> ⚔️\n\n")
		return dataString, nil
	case types.MUSIC:
		// db request
		dataString = fmt.Sprintf("🎵 <b>Плейлисты</b> 🎵\n\n")
		return dataString, nil
	case types.CHARACTER:
		// db request
		dataString = fmt.Sprintf("NYI\n\n")
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
