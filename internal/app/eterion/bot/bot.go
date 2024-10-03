package bot

import (
	"os"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

var keyboardMainMenu = tu.Keyboard(
	tu.KeyboardRow(
		tu.KeyboardButton("🌐 Кампании"),
		tu.KeyboardButton("🗺️ Карты"),
	),
	tu.KeyboardRow(
		tu.KeyboardButton("🧙‍♂️ Персонаж"),
	),
).WithResizeKeyboard()

func Init() {
	bot, err := telego.NewBot(os.Getenv("BOT_TOKEN"), telego.WithDefaultDebugLogger())
	if err != nil {
		panic("Unable to start bot:" + err.Error())
	}
}
