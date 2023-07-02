package models

import "time"

type Job struct {
	Logo           string    `json:"logo"`
	JobTitle       string    `json:"job_title"`
	Location       string    `json:"location"`
	Description    string    `json:"description"`
	HowToApply     string    `json:"how_to_apply"`
	Requirements   string    `json:"requirements"`
	Experience     string    `json:"experience"`
	Address        string    `json:"Address"`
	Categories     Category  `json:"Categories"`
	JobTypes       JobTypes  `json:"job_types"`
	Salary         float64   `json:"salary"`
	SubmissionDate time.Time `json:"submission_date"`
	Deadline       Date      `json:"deadline" binding:"required,futuredate"`
}

type JobTypes struct {
	JobTypes []string `json:"job_types" form:"job_types[]"`
}

type Category struct {
	Category []string `json:"category" form:"categories[]"`
}
