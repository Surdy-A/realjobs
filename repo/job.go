package repo

import (
	"database/sql"
	"fmt"
	"log"
	"realjobs/models"

	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
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

func CreateJob(j models.Job) error {
	ConnectToDB()
	insertStmt := `INSERT INTO job(logo, jobtitle, location, description, howtoApply, 
		requirements, experience, address, categories, jobtypes, salary, submissiondate, deadline) 
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12, $13)`

	_, err := db.Exec(insertStmt, j.Logo, j.JobTitle, j.Location, j.Description, j.HowToApply, 
		j.Requirements, j.Experience, j.Address, pq.Array(j.Categories.Categories), 
		pq.Array(j.JobTypes.JobTypes), j.Salary, j.SubmissionDate, j.Deadline)

	if err != nil {
		return err
	}

	return nil
}
