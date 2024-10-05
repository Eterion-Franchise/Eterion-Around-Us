package bot

import (
	"eterion_around_us/internal/app/eterion/database"
	"math/rand"

	"github.com/mymmrac/telego"
)

func UpdateSpamResponse(telegramUser telego.User) string {
	user := database.GetUserData(telegramUser.Username)

	roll := rand.Float64()

	if roll >= user.Variables.SpamResponseChance {
		// get secret phrase from db and return it
		database.UpdateUserData(database.User{
			UUID:          user.UUID,
			TgUserID:      user.TgUserID,
			IsWhitelisted: user.IsWhitelisted,
			IsGM:          user.IsGM,
			Variables: database.Variable{
				SpamResponseChance: database.GetSecretData().DefaultSpamResponseChance,
			},
		}, telegramUser.Username)
		return ""
	} else {
		database.UpdateUserData(database.User{
			UUID:          user.UUID,
			TgUserID:      user.TgUserID,
			IsWhitelisted: user.IsWhitelisted,
			IsGM:          user.IsGM,
			Variables: database.Variable{
				SpamResponseChance: increaseSpamResponceChance(user.Variables.SpamResponseChance),
			},
		}, telegramUser.Username)
	}

	return ""
}

func increaseSpamResponceChance(currentChance float64) float64 {
	return currentChance * database.GetSecretData().SpamResponseChanceFactor
}
