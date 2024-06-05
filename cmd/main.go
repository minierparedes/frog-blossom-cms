package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/reflection/frog_blossom_cms/api"
	"github.com/reflection/frog_blossom_cms/config"
	db "github.com/reflection/frog_blossom_cms/db/sqlc"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
