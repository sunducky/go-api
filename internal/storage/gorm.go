package storage

import (
	"fmt"
	"net/url"

	"github.com/sunducky/go-api-template/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BootsrapDB(env config.EnvVars) (*gorm.DB, error) {
	// Create the dsn from the env
	dsn := (&url.URL{
		Scheme:   env.DATABASE_SCHEME,
		User:     url.UserPassword(env.DATABASE_USERNAME, env.DATABASE_PASSWORD),
		Host:     fmt.Sprintf("%s:%d", env.DATABASE_HOST, env.DATABASE_PORT),
		Path:     env.DATABASE_NAME,
		RawQuery: "sslmode=disable",
	}).String()

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return DB, nil
}
