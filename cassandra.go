package config

import (
	"flag"
	"fmt"
)

var (
	cassandraConfig *CassandraConfig
)

//CassandraConfig holds Cassandra connections parameters
type CassandraConfig struct {
	URL      string
	Keyspace string
	Username string
	Password string
}

func (c *CassandraConfig) String() string {
	return fmt.Sprintf("CassandraConfig[URL=%v Keyspace=%v Username=%v Password=%v]", c.URL, c.Keyspace, c.Username, c.Password)
}

//BindCassandraConfiguration gets and binds, only if necessary, parameters for Cassandra connections
func BindCassandraConfiguration() *CassandraConfig {
	if cassandraConfig == nil {
		cassandraConfig = &CassandraConfig{}
		flag.StringVar(&cassandraConfig.URL, "cassandra_url", "127.0.0.1", "Cassandra url address")
		flag.StringVar(&cassandraConfig.Keyspace, "cassandra_keyspace", "fivecolors", "Cassandra keyspace")
		flag.StringVar(&cassandraConfig.Username, "cassandra_username", "fivecolors", "Cassandra username")
		flag.StringVar(&cassandraConfig.Password, "cassandra_password", "fivecolors", "Cassandra password")
	}
	return cassandraConfig
}
