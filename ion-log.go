package zerologr

// Disclamer
// This file contains functions that use actual zerologr, but make it compatible with pion/ion-log
//

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/go-logr/logr"
	"github.com/rs/zerolog"
)

const timeFormat = "2006-01-02 15:04:05.000"

var (
	log logr.Logger
	mu  sync.RWMutex
)

// Init creates new instances for logr.Logger
func Init(level string, fixByFile, fixByFunc []string) {
	zerologlvl := zerolog.GlobalLevel()
	switch level {
	case "trace":
		zerologlvl = zerolog.TraceLevel
	case "debug":
		zerologlvl = zerolog.DebugLevel
	case "info":
		zerologlvl = zerolog.InfoLevel
	case "warn":
		zerologlvl = zerolog.WarnLevel
	case "error":
		zerologlvl = zerolog.ErrorLevel
	}
	zerolog.TimeFieldFormat = timeFormat
	output := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false, TimeFormat: timeFormat}
	output.FormatTimestamp = func(i interface{}) string {
		return "[" + i.(string) + "]"
	}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[%-3s]", i))
	}
	output.FormatMessage = func(i interface{}) string {
		caller, file, line, _ := runtime.Caller(9)
		fileName := filepath.Base(file)
		funcName := strings.TrimPrefix(filepath.Ext((runtime.FuncForPC(caller).Name())), ".")
		var needfix bool
		for _, b := range fixByFile {
			if strings.Contains(fileName, b) {
				needfix = true
			}
		}
		for _, b := range fixByFunc {
			if strings.Contains(funcName, b) {
				needfix = true
			}
		}
		if needfix {
			caller, file, line, _ = runtime.Caller(8)
			fileName = filepath.Base(file)
			funcName = strings.TrimPrefix(filepath.Ext((runtime.FuncForPC(caller).Name())), ".")
		}
		return fmt.Sprintf("[%d][%s][%s] => %s", line, fileName, funcName, i)
	}

	l := zerolog.New(output).Level(zerologlvl).With().Timestamp().Logger()

	o := Options{
		Name:   "",
		Logger: &l,
	}

	mu.Lock()
	log = NewWithOptions(o)
	mu.Unlock()
}

// Config seems like isn't somewhere used, so just leave it here to don't let old code fail
type Config struct {
	Level string `mapstructure:"level"`
}

// Infof logs a formatted info level log to the console
func Infof(format string, v ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	log.Info(fmt.Sprintf(format, v...))

}

// Errorf logs a formatted error level log to the console
func Errorf(format string, v ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	log.Error(nil, fmt.Sprintf(format, v...))
}

// Panicf (Panic) is not support by logr interface, leave just an error
func Panicf(format string, v ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	log.Error(nil, fmt.Sprintf(format, v...))
}
