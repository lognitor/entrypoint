package ch

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/lognitor/entrypoint/configs"
)

type CH struct {
	config *configs.ClickHouseConfig
	Conn   driver.Conn
}

func NewClickHouse(conf *configs.ClickHouseConfig) (*CH, error) {
	ch := &CH{config: conf}

	if err := ch.connect(); err != nil {
		return nil, err
	}

	return ch, nil
}

func (c *CH) connect() error {
	if len(c.config.Hosts) < 1 {
		return fmt.Errorf("no hosts specified")
	}

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: c.config.Hosts,
		Auth: clickhouse.Auth{
			Database: c.config.Database,
			Username: c.config.User,
			Password: c.config.Password,
		},
	})

	if err != nil {
		return err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return err
	}

	c.Conn = conn
	return nil
}
