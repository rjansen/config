package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestSetupConfigSuccessfully(t *testing.T) {
	mockConfigPaths := []string{"./etc/config/config.yaml", "./etc/config/config.json"}
	for _, v := range mockConfigPaths {
		os.Args = []string{"config_utests", "-ecf", v}
		setupErr := Setup()
		assert.Nil(t, setupErr)
		assert.Equal(t, configFilePath, v)
		assert.Contains(t, viper.ConfigFileUsed(), v[1:])
		assert.True(t, configuration.InConfig("version"))
		assert.True(t, configuration.InConfig("type"))
		assert.Equal(t, configuration.GetString("type"), path.Ext(configFilePath)[1:])
	}
}
