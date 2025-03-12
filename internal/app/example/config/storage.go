package config

import (
	"time"

	"go-micro.dev/v4/config"
)

const defaultDBPort = 5432

type Storage struct {
	Book Postgres
}

type Postgres struct {
	Host            string
	Port            int
	Username        string
	Password        string
	DBName          string
	SSLMode         string
	ConnMaxLifetime time.Duration
	MaxOpenConns    int
	ConnMaxIdleTime time.Duration
	MaxIdleConns    int
}

func NewBookPostgres(cfg config.Config) Postgres {
	return Postgres{
		Host:            cfg.Get("db", "postgres", "book", "host").String("localhost"),
		Port:            cfg.Get("db", "postgres", "book", "port").Int(defaultDBPort),
		Username:        cfg.Get("db", "postgres", "book", "username").String("postgres"),
		Password:        cfg.Get("db", "postgres", "book", "password").String("password"),
		DBName:          cfg.Get("db", "postgres", "book", "dbname").String("postgres"),
		SSLMode:         cfg.Get("db", "postgres", "book", "sslmode").String("disable"),
		ConnMaxLifetime: cfg.Get("db", "postgres", "book", "connMaxLifetime").Duration(time.Second * 180),
		MaxOpenConns:    cfg.Get("db", "postgres", "book", "maxOpenConns").Int(10),
		ConnMaxIdleTime: cfg.Get("db", "postgres", "book", "connMaxIdleTime").Duration(time.Second * 240),
		MaxIdleConns:    cfg.Get("db", "postgres", "book", "maxIdleConns").Int(5),
	}
}

func newStorage(cfg config.Config) Storage {
	return Storage{
		Book: NewBookPostgres(cfg),
	}
}
