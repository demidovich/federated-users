package postgres

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SSLMode  bool
	Driver   string
}

func NewConnection(cfg Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Dbname,
		cfg.Password,
	)

	db, err := sqlx.Connect(cfg.Driver, dataSourceName)
	if err != nil {
		return nil, errors.New("Failed database connect: " + err.Error())
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New("Failed database ping: " + err.Error())
	}

	return db, nil
}

func NewConnectOrFail(cfg Config) *sqlx.DB {
	db, err := NewConnection(cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return db
}
