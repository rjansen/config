package config

import (
	"flag"
	"fmt"
	"github.com/vharitonsky/iniflags"
	"os"
	"time"
)

var (
	configuration *Configuration
)

type Configuration struct {
	Version string
	*HandlerConfig
	*ProxyConfig
	*CassandraConfig
	*DBConfig
	*SecurityConfig
	*CacheConfig
}

func (c *Configuration) String() string {
	return fmt.Sprintf("Configuration[Version[%v] ProxyConfig[%+v] DBConfig[%+v] SecurityConfig[%v]]", c.Version, c.ProxyConfig, c.DBConfig, c.SecurityConfig)
}

func BindConfiguration() *Configuration {
	if configuration == nil {
		configuration = &Configuration{}
		flag.StringVar(&configuration.Version, "version", fmt.Sprintf("transientbuild-%v", time.Now().UnixNano()), "Target bind address")

		configuration.HandlerConfig = BindHandlerConfiguration()
		configuration.ProxyConfig = BindProxyConfiguration()
		configuration.CassandraConfig = BindCassandraConfiguration()
		configuration.DBConfig = BindDBConfiguration()
		configuration.CacheConfig = BindCacheConfiguration()
	}
	return configuration
}

//Init initializes the flag system
func Init() {
	iniflags.Parse()
}

// Print error, usage and exit with code
func printErrorUsageAndExitWithCode(err string, code int) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	printUsage()
	os.Exit(code)
}

// Print command line help
func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [flags] [CONFIG]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nFlags:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nArguments:\n")
	fmt.Fprintf(os.Stderr, "  CONFIG: Config file path\n")
	fmt.Fprintf(os.Stderr, "\n")
}
