package configs

import (
	"errors"
	"os"
	"strings"
)

type ClickHouseConfig struct {
	Hosts    []string
	User     string
	Password string
	Database string
}

// NewClickHouseConfig create new config for clickhouse server
// IMPORTANT: host need port!
// Example: configs.NewClickHouseConfig([]string{"127.0.0.1:9001", "127.0.0.1:9002")}
func NewClickHouseConfig(hosts ...string) (*ClickHouseConfig, error) {
	user := os.Getenv("CLICKHOUSE_USER")
	if user == "" {
		return nil, errors.New("CLICKHOUSE_USER environment variable not set")
	}

	password := os.Getenv("CLICKHOUSE_PASSWORD")
	if password == "" {
		return nil, errors.New("CLICKHOUSE_PASSWORD environment variable not set")
	}

	database := os.Getenv("CLICKHOUSE_DATABASE")
	if database == "" {
		return nil, errors.New("CLICKHOUSE_DATABASE environment variable not set")
	}

	return &ClickHouseConfig{
		Hosts:    hosts,
		User:     user,
		Password: password,
		Database: database,
	}, nil
}

func (c *ClickHouseConfig) GetConnectionsString() string {
	return strings.Join(c.Hosts, ",")
}
