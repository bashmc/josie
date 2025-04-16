package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/topbash/split/handlers"
	"github.com/topbash/split/postgres"
	"github.com/topbash/split/services"
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
