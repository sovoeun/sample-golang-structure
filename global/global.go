package global

import (
	"example/sample/config"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"xorm.io/xorm"
)

var (
	Db     *xorm.Engine
	Config config.Server
	Viper  *viper.Viper
	Logger zerolog.Logger
)
