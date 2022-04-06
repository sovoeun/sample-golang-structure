package util

import (
	"example/sample/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"time"
)

func GetLogWriter(logConfig config.Log) (out io.Writer) {
	switch logConfig.Adapter {
	case "file":
		writer, err := rotatelogs.New(
			logConfig.Path+".%Y%m%d",
			rotatelogs.WithLinkName(logConfig.Path),
			rotatelogs.WithRotationCount(logConfig.ReverseDays),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			panic(err)
		}
		return writer
	}
	return os.Stdout
}
