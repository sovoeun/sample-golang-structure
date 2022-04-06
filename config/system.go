package config

type System struct {
	Env     string `mapstructure:"env"`
	Addr    string `mapstructure:"addr"`
	ShowSql bool   `mapstructure:"show-sql"`
	Sign    string `mapstructure:"sign"`
	Encrypt bool   `mapstructure:"encrypt"`
}
