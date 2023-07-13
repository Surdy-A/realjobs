package handlers

import (
	"fmt"
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
	price := c.PostForm("salary")
	deadline := c.PostForm("deadline")

	salary, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Conversion Error 1", err)
	}

	// var cat models.Categories
	// c.ShouldBind(&cat)

	var jobCategories models.Categories
	c.Bind(&jobCategories)

	fmt.Println(jobCategories.Categories)

	var job_types models.JobTypes
	c.ShouldBind(&job_types)

	var job models.Job
	c.ShouldBind(&job)
	layout := "2006-01-02"
	effectiveDate, err := time.Parse(layout, deadline)
	h := models.Job{
		Logo:           logo,
		JobTitle:       job_title,
		Location:       location,
		Description:    description,
		HowToApply:     how_to_apply,
		Requirements:   requirements,
		Experience:     experience,
		Address:        address,
		Categories:     jobCategories,
		JobTypes:       job_types,
		Salary:         salary,
		Deadline:       effectiveDate,
		SubmissionDate: time.Now(),
	}

	//var db *sql.DB
	err = repo.CreateJob(h)
	if err != nil {
		log.Fatal("Unable to create job:", err)
	}

	//var  []id
	//var  v
	// for id, v := range models.Jobs {
	// 	id= id
	// }

	c.HTML(http.StatusOK, "add_job.html", gin.H{
		"title": "Add Job",
	})

	c.Redirect(http.StatusPermanentRedirect, "/")
}

func AddJob(c *gin.Context) {

	c.HTML(http.StatusOK, "add_job.html", gin.H{
		"job_categories": models.JobCategories,
		"job_types_list": models.JobTypesList,
	})

	c.Redirect(http.StatusPermanentRedirect, c.Request.URL.Path)
}
