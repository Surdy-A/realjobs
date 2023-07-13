package repo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"realjobs/models"

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
		requirements, experience, address, categories, jobtype, salary, submissiondate, deadline) 
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12, $13)`

	_, err := db.Exec(insertStmt, &j.Logo, &j.JobTitle, &j.Location, &j.Description, &j.HowToApply,
		&j.Requirements, &j.Experience, &j.Address, pq.Array(&j.Categories.Categories),
		&j.JobType, &j.Salary, &j.SubmissionDate, &j.Deadline)

	if err != nil {
		return err
	}

	return nil
}

func GetJobs(j models.Job) ([]models.Job, error) {
	ConnectToDB()

	var jobs []models.Job
	selectStmt := `SELECT id, logo, jobtitle, location, description, howtoApply, 
	requirements, experience, address, categories, jobtype, salary, submissiondate, deadline
	FROM job`

	rows, err := db.Query(selectStmt)
	if err != nil {
		log.Fatal("Unable to retreive job: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&j.ID, &j.Logo, &j.JobTitle, &j.Location, &j.Description, &j.HowToApply,
			&j.Requirements, &j.Experience, &j.Address, pq.Array(&j.Categories.Categories),
			&j.JobType, &j.Salary, &j.SubmissionDate, &j.Deadline)

		if err != nil {
			fmt.Println("Unable to retreive job: ", err)
		}

		jobs = append(jobs, j)
	}

	return jobs, nil
}
