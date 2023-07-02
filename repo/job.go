package repo

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB
var err error

func ConnectToDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	var psqlInfo = user + "://" + user + ":" + password + "@" + host + "/" + dbname + "?sslmode=disable"
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("succesfuly connected to database")
}
