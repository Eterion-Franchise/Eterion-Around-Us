package types

type BotConfig struct {
	Debug           bool   `toml:"debug"`
	RequestCooldown int    `toml:"request_cooldown"` // NYI
	ParseMode       string `toml:"parse_mode"`
}

type WikiDataType int

const (
	CAMPAIGNS WikiDataType = iota
	MAPS      WikiDataType = iota
	BATTLES   WikiDataType = iota
	MUSIC     WikiDataType = iota
	CHARACTER WikiDataType = iota
)

type UserFlag int

const (
	FLAG_USER_WHITELISTED UserFlag = iota
	FLAG_USER_GM          UserFlag = iota
)
