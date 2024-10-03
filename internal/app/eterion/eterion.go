package eterion

import (
	"eterion_around_us/config"
	"eterion_around_us/internal/app/eterion/bot"
	"eterion_around_us/internal/app/eterion/database"
	"eterion_around_us/internal/app/eterion/telemetry"
)

func Start() {
	config.Init()
	database.Init()
	telemetry.Init()
	bot.Init()
}
