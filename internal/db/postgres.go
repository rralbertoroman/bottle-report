package db

import (
	"fmt"
	"log"

	"github.com/rralbertoroman/bottle-report/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Open(cfg config.Config) (db *gorm.DB){
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)
	
	db, err := 	gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return
}