package migrator

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"

	"go-micro-service-template/internal/app/migrator/config"
)

func Run(configPath string) error {
	cfg, err := config.New(configPath)
	if err != nil {
		return err
	}

	source := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v",
		cfg.Storage.Book.DBName,
		cfg.Storage.Book.Username,
		cfg.Storage.Book.Password,
		cfg.Storage.Book.Host,
		cfg.Storage.Book.Port,
		cfg.Storage.Book.SSLMode,
	)

	var db *sql.DB
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", source)
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}
		break
	}
	if err != nil {
		return err
	}
	defer db.Close()

	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}
