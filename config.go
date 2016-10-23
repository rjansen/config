package config

import (
	"flag"
	//"fmt"
	"github.com/matryer/resync"
	"strings"
	"time"
)

var (
	once    resync.Once
	ecfOnce resync.Once
)

// func init() {
// 	fmt.Println("config.init")
// }

//Get returns the singleton instance of the Configuration
func Get() Configuration {
	once.Do(func() {
		if configuration == nil {
			setupErr := setupViper()
			if setupErr != nil {
				panic(setupErr)
			}
			configuration = newViper()
		}
	})
	return configuration
}

//Setup initializes the package
func Setup() error {
	ecfOnce.Do(func() {
		if strings.TrimSpace(configFilePath) == "" {
			flag.StringVar(&configFilePath, "ecf", "", "The file configuration path")
		}
	})
	return setupViper()
}

func GetInterface(key string) interface{} {
	return Get().GetInterface(key)
}

func GetBool(key string) bool {
	return Get().GetBool(key)
}

func GetDuration(key string) time.Duration {
	return Get().GetDuration(key)
}

func GetFloat64(key string) float64 {
	return Get().GetFloat64(key)
}

func GetInt(key string) int {
	return Get().GetInt(key)
}

func GetInt64(key string) int64 {
	return Get().GetInt64(key)
}

func GetString(key string) string {
	return Get().GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return Get().GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return Get().GetStringMapString(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return Get().GetStringMapStringSlice(key)
}

func GetStringSlice(key string) []string {
	return Get().GetStringSlice(key)
}

func GetTime(key string) time.Time {
	return Get().GetTime(key)
}

func InConfig(key string) bool {
	return Get().InConfig(key)
}

func IsSet(key string) bool {
	return Get().IsSet(key)
}

func Unmarshal(rawVal interface{}) error {
	return Get().Unmarshal(rawVal)
}

func UnmarshalKey(key string, rawVal interface{}) error {
	return Get().UnmarshalKey(key, rawVal)
}
