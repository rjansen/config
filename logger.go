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
	File   string
	Level  int
	Format string
}

func (l *LoggerConfig) String() string {
	return fmt.Sprintf("LoggerConfig[File=%v Level=%v Foramt=%v]", l.File, l.Level, l.Format)
}

//BindLoggerConfiguration gets and binds, only if necessary, parameters for the application logger
func BindLoggerConfiguration() *LoggerConfig {
	if loggerConfig == nil {
		loggerConfig = &LoggerConfig{}
		flag.StringVar(&loggerConfig.File, "logger_file", "security.log", "Logger output file")
		flag.IntVar(&loggerConfig.Level, "logger_level", 5, "Logger output level")
		flag.StringVar(&loggerConfig.Format, "logger_format", "%{time:2006-01-02T15:04:05.999Z-07:00} %{id:03x} [%{level:.5s}] %{shortfunc} %{message}", "Logger output format")
	}
	return loggerConfig
}
