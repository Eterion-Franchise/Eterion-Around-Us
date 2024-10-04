package bot

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

var CampaignsButton = tu.KeyboardButton("📖 Игры")
var MapsButton = tu.KeyboardButton("🗺️ Карты")
var BattlesButton = tu.KeyboardButton("⚔️ Битвы")
var MusicButton = tu.KeyboardButton("🎵 Музыка")
var CharButton = tu.KeyboardButton("🧙‍♂️ Персонаж")

var KeyboardMainMenu = tu.Keyboard(
	tu.KeyboardRow(
		CampaignsButton,
		MapsButton,
		BattlesButton,
		MusicButton,
	),
	tu.KeyboardRow(
		CharButton,
	),
).WithResizeKeyboard()
