package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rralbertoroman/bottle-report/internal/config"
)


func Open(cfg config.Config) (db *sql.DB){
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := 	sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err:= db.Ping(); err != nil{
		log.Fatal(err)
	}

	return
}