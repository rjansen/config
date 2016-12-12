package migi

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
	"time"
)

var (
	testArgs = os.Args
)

func clean(t assert.TestingT) {
	reset(t)
}

func reset(t assert.TestingT) {
	configuration = nil
	once.Reset()
}

func TestSetupConfigSuccessfully(t *testing.T) {
	mockConfigPaths := []string{"./test/etc/config/config.yaml", "./test/etc/config/config.json"}
	for _, v := range mockConfigPaths {
		clean(t)
		os.Args = append(testArgs, "-ecf", v, "-edf")
		setupErr := Setup()
		assert.Nil(t, setupErr)
		assert.Equal(t, configFilePath, v)
		assert.Contains(t, viper.ConfigFileUsed(), v[1:])
		assert.True(t, Debug())
		assert.True(t, InConfig("version"))
		assert.True(t, InConfig("type"))
		assert.Equal(t, GetString("type"), path.Ext(configFilePath)[1:])
	}
}

func TestSetupConfigErr(t *testing.T) {
	cases := []struct {
		configFilePath string
		errContains    string
	}{
		{"/opt/filenotfound", "config.ErrReadConfigFile"},
		{"/opt/filenotfound.json", "config.ErrReadConfigFile"},
		{"", "config.ErrInvalidConfigFilePath"},
	}
	for _, v := range cases {
		clean(t)
		os.Args = append(testArgs, "-ecf", v.configFilePath)
		setupErr := Setup()
		assert.NotNil(t, setupErr)
		assert.Contains(t, setupErr.Error(), v.errContains)
	}
}

func TestGetConfigSuccessfully(t *testing.T) {
	cases := []struct {
		ecf        string
		configFile string
		configType string
	}{
		{"./test/etc/config/config.yaml", "config.yaml", "yaml"},
		{"./test/etc/config/config.json", "config.yaml", "yaml"},
		{"./test/etc/config/config.json", "config.yaml", "yaml"},
	}
	os.Args = append(os.Args, "-ecf", "")
	for _, v := range cases {
		os.Args = append(testArgs, "-ecf", v.ecf)
		cfg := Get()
		assert.NotNil(t, cfg)
		assert.Contains(t, viper.ConfigFileUsed(), v.configFile)
		assert.True(t, InConfig("type"))
		assert.Equal(t, GetString("type"), v.configType)

	}
}

func TestGetPanic(t *testing.T) {
	clean(t)
	os.Args = append(testArgs, "-ecf", "")
	assert.Panics(t, func() {
		Get()
	})
}

func TestGenericUnmarshalSuccessfully(t *testing.T) {
	clean(t)
	os.Args = append(testArgs, "-ecf", "./test/etc/config/unmarshal.yaml")
	cfg := Get()
	assert.NotNil(t, cfg)
	configMap := make(map[string]interface{})
	var err error
	err = cfg.Unmarshal(&configMap)
	assert.Nil(t, err)

	compareConfigMap := map[string]interface{}{
		"version": "unmarshal",
		"type":    "yaml",
		"subconfig": map[string]interface{}{
			"id":           "identifier",
			"ttl":          1024,
			"ispersistent": false,
		},
	}
	assert.EqualValues(t, configMap, compareConfigMap)
}

func TestStructUnmarshalSuccessfully(t *testing.T) {
	config := struct {
		Version    string `json:"version" mapstructure:"version"`
		ConfigType string `json:"type" mapstructure:"type"`
		Sub        struct {
			ID         string `json:"id" mapstructure:"id"`
			TTL        int    `json:"ttl" mapstructure:"ttl"`
			Persistent bool   `json:"isPersistent" mapstructure:"isPersistent"`
		} `json:"subConfig" mapstructure:"subConfig"`
	}{}
	clean(t)
	os.Args = append(testArgs, "-ecf", "./test/etc/config/unmarshal.yaml")
	var err error
	err = Unmarshal(&config)
	assert.Nil(t, err)

	assert.NotZero(t, config)
	assert.NotZero(t, config.Version)
	assert.NotZero(t, config.ConfigType)
	assert.NotZero(t, config.Sub)
	assert.NotZero(t, config.Sub.ID)
	assert.NotZero(t, config.Sub.TTL)
	assert.False(t, config.Sub.Persistent)
}

func TestStructUnmarshalVarSuccessfully(t *testing.T) {
	type Cfg struct {
		Version    string `json:"version" mapstructure:"version"`
		ConfigType string `json:"type" mapstructure:"type"`
		Sub        struct {
			ID         string `json:"id" mapstructure:"id"`
			TTL        int    `json:"ttl" mapstructure:"ttl"`
			Persistent bool   `json:"isPersistent" mapstructure:"isPersistent"`
		} `json:"subConfig" mapstructure:"subConfig"`
	}
	clean(t)
	os.Args = append(testArgs, "-ecf", "./test/etc/config/unmarshal.yaml")
	var err error
	var config *Cfg
	err = Unmarshal(&config)
	assert.Nil(t, err)

	assert.NotZero(t, config)
	assert.NotZero(t, config.Version)
	assert.NotZero(t, config.ConfigType)
	assert.NotZero(t, config.Sub)
	assert.NotZero(t, config.Sub.ID)
	assert.NotZero(t, config.Sub.TTL)
	assert.False(t, config.Sub.Persistent)
}

func TestUnmarshalKeySuccessfully(t *testing.T) {
	config := struct {
		ID         string `json:"id" mapstructure:"id"`
		TTL        int    `json:"ttl" mapstructure:"ttl"`
		Persistent bool   `json:"isPersistent" mapstructure:"isPersistent"`
	}{}
	clean(t)
	os.Args = append(testArgs, "-ecf", "./test/etc/config/unmarshal.yaml")
	var err error
	err = UnmarshalKey("subConfig", &config)
	assert.Nil(t, err)

	assert.NotZero(t, config)
	assert.NotZero(t, config.ID)
	assert.NotZero(t, config.TTL)
	assert.False(t, config.Persistent)
}

func TestGetKeysSuccessfully(t *testing.T) {
	clean(t)
	os.Args = append(testArgs, "-ecf", "./test/etc/config/alltypes.yaml")
	assert.Equal(t, GetInterface("version"), "alltypes")
	assert.Equal(t, GetBool("boolKey"), true)
	assert.Equal(t, GetDuration("durationKey"), (time.Second*10)+time.Millisecond)
	assert.Equal(t, GetFloat64("floatKey"), float64(0.64))
	assert.Equal(t, GetInt("intKey"), 4096)
	assert.Equal(t, GetInt64("intKey"), int64(4096))
	assert.Equal(t, GetString("stringKey"), "lorem ipsum str")
	assert.NotEmpty(t, GetStringMap("subConfig"))
	assert.NotEmpty(t, GetStringMapString("subConfigString"))
	assert.NotEmpty(t, GetStringMapStringSlice("subConfigStringSlice"))
	assert.Equal(t, GetStringSlice("sliceKey"), []string{"slice1", "slice2", "slice3", "slice4"})
	assert.NotZero(t, GetTime("timeKey"))
	assert.True(t, IsSet("type"))
}
