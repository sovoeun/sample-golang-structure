package initialize

import (
	"example/sample/global"
	"github.com/spf13/viper"
)

func initViper() *viper.Viper {
	path := "./config.yaml"
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
	return v
}
