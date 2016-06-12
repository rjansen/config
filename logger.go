package config

import (
	"flag"
	"fmt"
)

var (
	loggerConfig *LoggerConfig
)

//LoggerConfig holds log parameters
type LoggerConfig struct {
	LoggerFile  string
	LoggerLevel int
}

func (l *LoggerConfig) String() string {
	return fmt.Sprintf("LoggerConfig[LoggerFile=%v]", l.LoggerFile)
}

//BindLoggerConfiguration gets and binds, only if necessary, parameters for the application logger
func BindLoggerConfiguration() *LoggerConfig {
	if loggerConfig == nil {
		loggerConfig = &LoggerConfig{}
		flag.StringVar(&loggerConfig.LoggerFile, "logger_file", "epedion.log", "Logger output file")
		flag.IntVar(&loggerConfig.LoggerLevel, "logger_level", 3, "Logger output level")
	}
	return loggerConfig
}
