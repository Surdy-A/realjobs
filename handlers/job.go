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
	address := c.PostForm("address")
	price := c.PostForm("salary")
	deadline := c.PostForm("deadline")
	job_type := c.PostForm("job_type")
	salary, err := strconv.ParseFloat(price, 64)

	if err != nil {
		log.Println("Conversion Error:", err)
	}

	job_experience := c.PostForm("experience")
	experience, err := strconv.ParseInt(job_experience, 10, 64)

	if err != nil {
		log.Println("Conversion Error:", err)
	}

	layout := "2006-01-02"
	effectiveDate, err := time.Parse(layout, deadline)

	if err != nil {
		log.Println("Conversion Error:", err)
	}

	var jobCategories models.Categories
	c.Bind(&jobCategories)

	var job models.Job
	c.ShouldBind(&job)

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
		JobType:        job_type,
		Salary:         salary,
		Deadline:       effectiveDate,
		SubmissionDate: time.Now(),
	}

	err = repo.CreateJob(h)
	if err != nil {
		log.Fatal("Unable to create job:", err)
	}

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

func GetJobs(c *gin.Context) {
	var j models.Job
	jobs, err := repo.GetJobs(j)

	if err != nil {
		log.Fatal("Unable to get jobs:", err)
	}

	c.HTML(http.StatusOK, "job_lists.html", gin.H{
		"title": "Real Jobs",
		"jobs":  jobs,
	})
}

func GetJob(c *gin.Context) {
	id := c.Param("id")
	var j models.Job
	job, err := repo.GetJob(id, j)

	if err != nil {
		log.Fatal("Unable to get jobs:", err)
	}

	c.HTML(http.StatusOK, "job_detail.html", gin.H{
		"title": "Get Job" + id,
		"job":   job,
	})
}
