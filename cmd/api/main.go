package main

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/shcmd/split/handlers"
	"github.com/shcmd/split/mail"
	"github.com/shcmd/split/postgres"
	"github.com/shcmd/split/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logFatal("failed to load environment variables", err)
	}

	cfg := loadConfig()

	db, err := pgxpool.New(context.Background(), cfg.PostgresURL)
	if err != nil {
		logFatal("failed to connect to database", err)
	}
	defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		logFatal("failed to ping db", err)
	}

	mailer := mail.NewMailer(cfg.MailConfig)
	handler := handlers.NewHandler(services.NewUserService(postgres.NewUserStore(db), mailer))

	srv := newserver(handler)

	slog.Info("Starting server")
	err = srv.start()
	if err != nil {
		logFatal("failed to start server", err)
	}
}
