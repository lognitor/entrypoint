package main

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lognitor/entrypoint/configs"
	"log"
	"os"
	"strings"
)

func main() {
	chHosts := strings.Split(os.Getenv("CLICKHOUSE_HOSTS"), ",")
	cfg, err := configs.NewClickHouseConfig(chHosts...)
	if err != nil {
		log.Fatalf("error loading clickhouse config: %v", err)
	}

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("clickhouse://%s?username=%s&password=%s&database=%s&x-multi-statement=true",
			cfg.Hosts[0], cfg.User, cfg.Password, cfg.Database))
	if err != nil {
		log.Fatalf("failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("error up migrations: %v", err)
	}

	log.Println("migrations applied")
}
