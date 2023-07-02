package main

import (
	"fmt"
	"net/http"
	"realjobs/models"
	"time"

	"github.com/gin-gonic/gin"
)
const (
    DDMMYYYY = "02/01/2006"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/", func(c *gin.Context) {
		//name := c.Param("name")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Real Jobs",
		})
	})

	router.GET("/jobs", func(c *gin.Context) {
		//name := c.Param("name")
		layout := "2006-01-02"
		job_deadline := "2023-11-04"
		deadline, err := time.Parse(layout, job_deadline)
		if err != nil {
			return
		}
		//date := time.Now().Format(DDMMYYYY)

		job := models.Job{Logo: "", JobTitle: "Digital Marketing Executive", Location: "Sacramento, California", JobTypes: models.JobTypes{},
			Deadline: models.Date(deadline), Description: "Description", HowToApply: "How to Apply", Requirements: "Requirements", Experience: "3",
			Address:    "Demo Address #8901 Marmora Road Chi Minh City, Vietnam",
			Categories: models.Category{},
			Salary:     600.700,
			SubmissionDate: time.Now(),
		}

		//fmt.Println(job)
		fmt.Println(deadline)

		c.HTML(http.StatusOK, "job_lists.html", gin.H{
			"title": "Real Jobs",
			"job":   job,
		})
	})

	router.Run(":8080")
}
