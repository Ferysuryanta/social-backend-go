package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("❌ Failed to parse database URL: %v", err)
	}

	// tuning untuk concurrency
	cfg.MaxConns = 20
	cfg.MinConns = 5
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.HealthCheckPeriod = 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DB, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatal("create pool:", err)
	}

	if err := DB.Ping(ctx); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	log.Println("Connected to database")
}
