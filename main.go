package main

import (
	"realjobs/handlers"

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
	router.GET("/", handlers.Home)
	router.GET("/add_job", handlers.AddJob)
	router.POST("/add_job", handlers.CreateJob)
	router.GET("/jobs", handlers.GetJobs)
	router.GET("/job_:id", handlers.GetJob)

	router.Run(":8081")
}
