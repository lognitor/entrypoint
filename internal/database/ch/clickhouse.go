package ch

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/lognitor/entrypoint/configs"
	"time"
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
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		//Compression: &clickhouse.Compression{
		//	Method: clickhouse.CompressionLZ4,
		//},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
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
