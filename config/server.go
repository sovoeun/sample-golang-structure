package config

type Server struct {
	Db     Db     `mapstructure:"db"`
	System System `mapstructure:"system"`
	Log    Log    `mapstructure:"log"`
}
