package eterion

import (
	"eterion_around_us/config"
	"eterion_around_us/internal/app/eterion/bot"
)

func Start() {
	config.Init()
	//database.Init()
	bot.Init()
}
