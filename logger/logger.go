package logger

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLog(serviceName string) {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short
		return file + ":" + strconv.Itoa(line)
	}
	Log = zerolog.New(os.Stdout).With().Timestamp().Caller().Str("service-name", serviceName).Logger()
}
