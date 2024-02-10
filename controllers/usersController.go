package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var body struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to read request body",
		}) 
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to hash the password",
		})
		return
	}
	user := models.User{
		FirstName: body.FirstName,
		LastName: body.LastName,
		Password: string(hash),
		Username: body.Username, 
	}
	
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{ //201
		"message": "User created",
	}) 
}

func GetUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	type publicUser struct {
		ID             string    `json:"id"`
		FirstName      string    `json:"first_name"`
		LastName       string    `json:"last_name"`
		Username       string    `json:"username"`
		AccountCreated time.Time `json:"account_created"`
		AccountUpdated time.Time `json:"account_updated"`
	}

	public := publicUser{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		AccountCreated: user.AccountCreated,
		AccountUpdated: user.AccountUpdated,
	}
	c.JSON(http.StatusOK, public) //200
}

func UpdateUser(c *gin.Context) {
	user := c.MustGet("user").(models.User)

  var updated models.User

  if c.ShouldBindJSON(&updated) != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Failed to read request body",
    })
    return
  }

	// Check disallowed fields in the request body
  if updated.ID != "" ||  updated.Username != "" || updated.AccountCreated != (time.Time{}) || updated.AccountUpdated != (time.Time{}) {
    c.JSON(http.StatusBadRequest, gin.H{ //400
      "error": "Attempt to update disallowed field",
    })
    return
  }

  // Return 204 if no changes were made
	if updated.FirstName == "" && updated.LastName == "" && updated.Password == "" {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No changes were made",
		})
		return
	}

	user.FirstName = updated.FirstName
	user.LastName = updated.LastName
	if updated.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(updated.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash the password",
			})
			return
		}
		user.Password = string(hash)
	}

	user.AccountUpdated = time.Now()
  result := initializers.DB.Save(&user)
  if result.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Failed to update user",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "User updated successfully": user,
	})
}