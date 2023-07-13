package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Job struct {
	Logo           string     `json:"logo"`
	JobTitle       string     `json:"job_title"`
	Location       string     `json:"location"`
	Description    string     `json:"description"`
	HowToApply     string     `json:"how_to_apply"`
	Requirements   string     `json:"requirements"`
	Experience     string     `json:"experience"`
	Address        string     `json:"Address"`
	Categories     Categories `json:"Categories"`
	JobTypes       JobTypes   `json:"job_types"`
	Salary         float64    `json:"salary"`
	SubmissionDate time.Time  `json:"submission_date"`
	Deadline       time.Time       `json:"deadline" binding:"required"`
}

type JobTypes struct {
	JobTypes []string `json:"job_types" form:"job_types[]"`
}
type Categories struct {
	Categories []string `json:"categories" form:"categories[]"`
}

var JobCategories = map[int]string{1: "Fashion",
	2:  "Construction",
	3:  "Management",
	4:  "Sales",
	5:  "Human Resources",
	6:  "Marketing",
	7:  "Computer Programmer",
	8:  "Engineer",
	9:  "Engineering",
	10: "Accountant",
	11: "Technician",
	12: "Computers and information technology",
	13: "Electrician",
	14: "Administrative assistant",
	15: "Teacher",
	16: "Customer Service",
	17: "Law",
	18: "Accounting",
	19: "Police officer",
	20: "Photographer",
	21: "Sales Management",
	22: "Architect",
	23: "Project management",
	24: "Data science",
	25: "Frontend Developer",
	26: "Backend Developer",
	27: "Fullstack Developer",
	28: "Data Engineer",
	29: "Data Analyst",
}

var JobTypesList = map[int]string{
	1: "Onsite",
	2: "Remote",
	3: "Hybrid",
	4: "Fulltime",
	5: "Contract",
	6: "Parttime",
}

// fmt.Println()
func (h Job) Value() (driver.Value, error) {
	return json.Marshal(h)
}

func (h *Job) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &h)
}
