package bot

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

var CampaignsButton = tu.KeyboardButton("📖 Игры")
var MapsButton = tu.KeyboardButton("🗺️ Карты")
var BattlesButton = tu.KeyboardButton("⚔️ Битвы")
var MusicButton = tu.KeyboardButton("🎵 Музыка")
var CharButton = tu.KeyboardButton("🧙‍♂️ Персонаж")

var LookAroundButton = tu.KeyboardButton("Осмотреться")

var MoveForward = tu.KeyboardButton("Пройти дальше")
var TurnBackButton = tu.KeyboardButton("Повернуть назад")

var TouchButton = tu.KeyboardButton("Прикоснуться к обелиску")

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

var KeyboardEntryPhase1 = tu.Keyboard(
	tu.KeyboardRow(
		LookAroundButton,
	),
)

var KeyboardEntryPhase2 = tu.Keyboard(
	tu.KeyboardRow(
		MoveForward,
	),
	tu.KeyboardRow(
		TurnBackButton,
	),
)

var KeyboardEntryPhase3 = tu.Keyboard(
	tu.KeyboardRow(
		TouchButton,
	),
)
