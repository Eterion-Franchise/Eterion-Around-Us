package types

type BotConfig struct {
	Debug            bool `toml:"debug"`
	RequestCooldown  int  `toml:"request_cooldown"`
	CollectTelemetry bool `toml:"collect_telemetry"`
}
