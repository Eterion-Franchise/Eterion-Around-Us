package bot

import (
	"log"
	"os"

	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

func Init() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		panic("Unable to start bot:" + err.Error())
	}
	defer bot.StopLongPolling()

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := telegohandler.NewBotHandler(bot, updates)

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		_, err := bot.SendMessage(&telego.SendMessageParams{
			ChatID:      message.Chat.ChatID(),
			Text:        "<i>Вы в архиве Этериона. Какое знание вы хотите открыть сегодня?</i>",
			ParseMode:   "html",
			ReplyMarkup: KeyboardMainMenu,
		})
		if err != nil {
			log.Print(err)
		}
	})

	bh.Start()
}
