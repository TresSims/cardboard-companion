package db

import (
	"cardboard-companion/config"
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func init() {
	var dialect gorm.Dialector

	config.InitConfig()

	conf := config.Load()

	parsed, err := url.Parse(conf.DatabaseURL)
	if err != nil {
		panic(err)
	}

	switch parsed.Scheme {
	case "sqlite":
		dialect = sqlite.Open(parsed.Host)
	case "postgresql", "postgres", "psql":
		password, passwordSet := parsed.User.Password()
		if !passwordSet {
			panic("No password set on connection URI")
		}

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			parsed.Hostname(),
			parsed.User.Username(),
			password,
			parsed.Path[1:],
			parsed.Port(),
		)

		dialect = postgres.Open(dsn)
	}

	Connection, _ = gorm.Open(dialect)
}
