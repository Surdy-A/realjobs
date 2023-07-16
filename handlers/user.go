package handlers

import (
	"log"
	"net/http"
	"realjobs/models"
	"realjobs/repo"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User

	c.ShouldBind(&user)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	if err != nil {
		log.Fatal("Unable to generate hash: ", err)
	}


	u := models.User{
		FirstName:    first_name,
		LastName:     last_name,
		Email:        email,
		Password:     password,
		PasswordHash: string(passwordHash),
	}

	err = repo.CreateUser(u)

	if err != nil {
		log.Fatal("Unable to create user:", err)
	}

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}
 
func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
 