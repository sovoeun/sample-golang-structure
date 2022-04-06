package initialize

import (
	"example/sample/global"
	"github.com/rs/zerolog"
	"time"
)

func InitComponents() {
	zerolog.TimeFieldFormat = time.RFC3339
	global.Viper = initViper()
	global.Db = initDb()
}
