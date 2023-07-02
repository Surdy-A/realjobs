package handlers

import (
	"log"
	"net/http"
	"realjobs/models"
	"realjobs/repo"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {
	logo := c.PostForm("logo")
	job_title := c.PostForm("job_title")
	location := c.PostForm("location")
	description := c.PostForm("description")
	how_to_apply := c.PostForm("how_to_apply")
	requirements := c.PostForm("requirements")
	experience := c.PostForm("experience")
	address := c.PostForm("address")
	price := c.PostForm("price")
	//	deadline := c.PostForm("deadline")

	salary, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Conversion Error 1", err)
	}
	var cat models.Category
	c.ShouldBind(&cat)

	var job_types models.JobTypes
	c.ShouldBind(&job_types)

	var job models.Job
	c.ShouldBind(&job)

	h := models.Job{
		Logo:         logo,
		JobTitle:     job_title,
		Location:     location,
		Description:  description,
		HowToApply:   how_to_apply,
		Requirements: requirements,
		Experience:   experience,
		Address:      address,
		Categories:   cat,
		JobTypes:     job_types,
		Salary:       salary,
		//Deadline:       models.Date{deadline},
		SubmissionDate: time.Now(),
	}

	//var db *sql.DB
	err = repo.CreateJob(h)
	if err != nil {
		log.Fatal("Unable to create job:", err)
	}

	c.HTML(http.StatusOK, "add_job.tmpl", gin.H{
		"title": "Add Job",
	})

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
}

func AddJob(c *gin.Context) {

	c.HTML(http.StatusOK, "add_job.html", gin.H{})

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
}
