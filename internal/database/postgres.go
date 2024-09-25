package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arifinoid/room-reservation-api/internal/config"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
		return nil, err
	}
	log.Println("connected to database")
	return db, nil
}
