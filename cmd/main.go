package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/reflection/frog-blossom-cms/api"
	"github.com/reflection/frog-blossom-cms/config"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
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
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = api.Start(server, config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
