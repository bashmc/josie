package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xpmc/split/handlers"
	"github.com/xpmc/split/postgres"
	"github.com/xpmc/split/services"
)

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	handler := handlers.NewAppHandler(services.NewUserService(postgres.NewUserStore(db)))

	srv := NewServer(handler)

	log.Println("Starting server....")
	log.Fatal(srv.start())
}
