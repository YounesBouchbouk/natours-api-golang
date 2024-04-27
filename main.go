package main

import (
	"database/sql"
	"log"

	"github.com/YounesBouchbouk/natours-api-golang/api"
	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/natours?sslmode=disable"
)

func main() {
	// Open a database value
	conn, err := sql.Open(dbDriver, dbSource)
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
	server := api.NewServer(store)

	// Start the server
	err = server.Start(":3001")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
