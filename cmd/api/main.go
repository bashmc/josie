package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gitkobie/split/handlers"
	"github.com/gitkobie/split/postgres"
	"github.com/gitkobie/split/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logFatal("failed to load environment variables", err)
	}

	db, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		logFatal("failed to connect to database", err)
	}
	defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		logFatal("failed to ping db", err)
	}

	handler := handlers.NewAppHandler(services.NewUserService(postgres.NewUserStore(db)))

	srv := NewServer(handler)

	slog.Info("Starting server")
	err = srv.start()
	if err != nil {
		logFatal("failed to start server", err)
	}
}
