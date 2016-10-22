package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	//path "path/filepath"
	"strings"
	"time"
)

type viperConfiguration struct {
	configFilePath string
	viper          *viper.Viper
}

func (v viperConfiguration) String() string {
	return fmt.Sprintf("viperConfiguration configFilePath=%s viperIsNil=%t", v.configFilePath, v.viper == nil)
}

func (v viperConfiguration) GetInterface(key string) interface{} {
	return v.viper.Get(key)
}

func (v viperConfiguration) GetBool(key string) bool {
	return v.viper.GetBool(key)
}

func (v viperConfiguration) GetDuration(key string) time.Duration {
	return v.viper.GetDuration(key)
}

func (v viperConfiguration) GetFloat64(key string) float64 {
	return v.viper.GetFloat64(key)
}

func (v viperConfiguration) GetInt(key string) int {
	return v.viper.GetInt(key)
}

func (v viperConfiguration) GetInt64(key string) int64 {
	return v.viper.GetInt64(key)
}

func (v viperConfiguration) GetString(key string) string {
	return v.viper.GetString(key)
}

func (v viperConfiguration) GetStringMap(key string) map[string]interface{} {
	return v.viper.GetStringMap(key)
}

func (v viperConfiguration) GetStringMapString(key string) map[string]string {
	return v.viper.GetStringMapString(key)
}

func (v viperConfiguration) GetStringMapStringSlice(key string) map[string][]string {
	return v.viper.GetStringMapStringSlice(key)
}

func (v viperConfiguration) GetStringSlice(key string) []string {
	return v.viper.GetStringSlice(key)
}

func (v viperConfiguration) GetTime(key string) time.Time {
	return v.viper.GetTime(key)
}

func (v viperConfiguration) InConfig(key string) bool {
	return v.viper.InConfig(key)
}

func (v viperConfiguration) IsSet(key string) bool {
	return v.viper.IsSet(key)
}

//func (v viperConfiguration) Reset()
func (v viperConfiguration) Unmarshal(rawVal interface{}) error {
	return v.viper.Unmarshal(rawVal)

}
func (v viperConfiguration) UnmarshalKey(key string, rawVal interface{}) error {
	return v.viper.UnmarshalKey(key, rawVal)
}

func setupViper() error {
	flag.Parse()
	if strings.TrimSpace(configFilePath) == "" {
		fmt.Println("comfig.setupViper.ErrInvalidConfigFilePath ", configFilePath)
		return fmt.Errorf("config.ErrInvalidConfigFilePath ecf=%v\n", configFilePath)
	}
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return fmt.Errorf("config.ErrReadConfigFile ecf=%v message='%s'\n", configFilePath, err)
	}
	configuration = &viperConfiguration{
		configFilePath: configFilePath,
		viper:          viper.GetViper(),
	}
	fmt.Printf("config.settedViper configuration=%s\n", configuration)
	return nil
}
