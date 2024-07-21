package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/jackc/pgx/v4/pgxpool"
)



func main() {

	dbUri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_DBUSER"),
		os.Getenv("POSTGRES_DBPASS"),
		os.Getenv("POSTGRES_DBHOST"),
		os.Getenv("POSTGRES_DBPORT"),
		os.Getenv("POSTGRES_DBNAME_GO"),
	)
	fmt.Println(dbUri)
	db, err := pgxpool.Connect(context.Background(), dbUri)
	// check any err in connection
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected via pgxPool")

}
