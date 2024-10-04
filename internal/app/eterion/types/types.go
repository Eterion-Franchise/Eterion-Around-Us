package types

type BotConfig struct {
	Debug           bool `toml:"debug"`
	RequestCooldown int  `toml:"request_cooldown"` // NYI
}

type WikiDataType int

const (
	CAMPAIGNS WikiDataType = iota
	MAPS      WikiDataType = iota
	BATTLES   WikiDataType = iota
	MUSIC     WikiDataType = iota
)
