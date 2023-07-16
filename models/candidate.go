package models

import "time"

type Candidate struct {
	ID             string    `json:"id"`
	Logo           string    `json:"logo"`
	Name           string    `json:"name"`
	Location       string    `json:"location"`
	CV             string    `json:"cv"`
	CoverLetter    string    `json:"cover_letter"`
	Experience     int64     `json:"experience"`
	Address        string    `json:"Address"`
	Skills         Skills    `json:"Categories"`
	JobType        string    `json:"job_types"`
	HourRate       float64   `json:"hour_rate"`
	SubmissionDate time.Time `json:"submission_date"`
}

type Skills struct {
	Skills []string `json:"skills" form:"skills[]"`
}
