package config

import (
	"eterion_around_us/internal/app/eterion/types"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pelletier/go-toml/v2"
)

var BotConfig types.BotConfig

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var data types.BotConfig
	content, err := os.ReadFile("./config/config.toml")
	if err != nil {
		panic("Unable to read config file:" + err.Error())
	}

	err = toml.Unmarshal(content, &data)
	if err != nil {
		panic("Unable to unmarshal config file:" + err.Error())
	}

	BotConfig = data
}
