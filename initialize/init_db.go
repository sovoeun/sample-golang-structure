package initialize

import (
	"example/sample/global"
	logadaptor "example/sample/initialize/log-adaptor"
	"example/sample/util"
	"fmt"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func initDb() *xorm.Engine {
	dbConfig := global.Config.Db
	if len(dbConfig.Host) == 0 {
		global.Logger.Warn().Msg("Database parameters are not configured")
		return nil
	}
	password := dbConfig.Password

	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbConfig.User, password, dbConfig.DbName, dbConfig.Host, dbConfig.Port)
	engine, err := xorm.NewEngine("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = engine.Ping(); err != nil {
		panic(err)
	}
	writer := util.GetLogWriter(dbConfig.Log)
	xormLogger := logadaptor.NewXormLogger(writer)
	xormLogger.SetLevel2(dbConfig.Log.Level)
	if dbConfig.MaxIdleConn > 0 {
		engine.SetMaxIdleConns(dbConfig.MaxIdleConn)
	}
	engine.SetLogger(xormLogger)
	engine.EnableSessionID(global.Config.System.ShowSql)
	engine.ShowSQL(global.Config.System.ShowSql)
	return engine
}
