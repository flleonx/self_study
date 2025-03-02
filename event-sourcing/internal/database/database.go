package database

import (
	"event-sourcing/configuration"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewDatabase(config configuration.Config) (*sqlx.DB, error) {
	return sqlx.Connect(
		"pgx",
		fmt.Sprintf(
			"postgresql://%s:%s@%s/%s?sslmode=disable",
			config.User,
			config.Password,
			config.Host,
			config.Name,
		),
	)
}
