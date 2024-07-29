package main

import (
	"database/sql"
	"izumarth_grpc/api"
	db "izumarth_grpc/db/sqlc"
	"izumarth_grpc/db/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfing(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
