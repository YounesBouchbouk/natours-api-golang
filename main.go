package main

import (
	"database/sql"
	"log"

	"github.com/YounesBouchbouk/natours-api-golang/api"
	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	"github.com/YounesBouchbouk/natours-api-golang/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open a database value
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer conn.Close()

	// Ping the database to check if the connection is working
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping database:", err)
	}

	// Create a new store
	store := db.NewStore(conn)

	// Create a new server
	server, _ := api.NewServer(store, &config)

	// Start the server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
