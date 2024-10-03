package telemetry

import (
	"eterion_around_us/config"
	"log"
)

func Init() {
	if !config.BotConfig.CollectTelemetry {
		log.Println("Telemetry collection is turned off. Skipping initialization...")
		return
	}
}
