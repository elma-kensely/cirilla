package types

//Config structure to store configuration datas
//nolint:lll
type Config struct {
	TelegramToken string `envconfig:"CIRILLA_TOKEN"`
	Debug         bool   `envconfig:"CIRILLA_DEBUG" default:"false"`
	Timeout       int    `envconfig:"CIRILLA_TIMEOUT" default:"86400"`
	CommandPrefix string `envconfig:"CIRILLA_CMD_PREFIX" default:"/"`
}
