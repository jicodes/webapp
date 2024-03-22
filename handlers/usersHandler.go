package handlers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"

	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/models"
)

func CreateUser(c *gin.Context) {
	logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if logErr != nil {
		panic(logErr)
	}
	defer logFile.Close()
	logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

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

	if !models.ValidateEmail(body.Username) {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Username must be a valid email",
		})

		logger.Error().Msg("Username must be a valid email")
		return
	}

	var existingUser models.User
	err := initializers.DB.First(&existingUser, "username = ?", body.Username).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "User already exists",
		})

		logger.Error().Msg("User already exists")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to hash the password",
		})
		return
	}
	
	newUser := models.User{
		FirstName: body.FirstName,
		LastName: body.LastName,
		Password: string(hash),
		Username: strings.ToLower(body.Username),
	}

	result := initializers.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{ //400
			"error": "Failed to create user",
		})
		return
	}

	publicUser := models.PublicUser{
		ID:             newUser.ID,
		FirstName:      newUser.FirstName,
		LastName:       newUser.LastName,
		Username:       newUser.Username,
		AccountCreated: newUser.AccountCreated,
		AccountUpdated: newUser.AccountUpdated,
	}

	c.JSON(http.StatusCreated, publicUser)
	logger.Info().Msg("User created successfully") 
}

func GetUser(c *gin.Context) {
	logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if logErr != nil {
		panic(logErr)
	}
	defer logFile.Close()
	logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

	user := c.MustGet("user").(models.User)
	public := models.PublicUser{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		AccountCreated: user.AccountCreated,
		AccountUpdated: user.AccountUpdated,
	}
	c.JSON(http.StatusOK, public) //200
	logger.Info().Msg("User retrieved successfully")
}

func UpdateUser(c *gin.Context) {
	logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if logErr != nil {
		panic(logErr)
	}
	defer logFile.Close()
	logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

	user := c.MustGet("user").(models.User)
  var updated models.User

  if c.ShouldBindJSON(&updated) != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Request body should be  be in JSON format",
    })
		logger.Error().Msg("Request body should be in JSON format")
    return
  }

	// Check disallowed fields in the request body
  if updated.ID != "" ||  updated.Username != "" || updated.AccountCreated != (time.Time{}) || updated.AccountUpdated != (time.Time{}) {
    c.JSON(http.StatusBadRequest, gin.H{ //400
      "error": "You can only update the fields of FirstName, LastName and Password",
    })
		logger.Error().Msg("Disallowed fields in the request body")
    return
  }

  // Return 204 if no changes were made
	if updated.FirstName == "" && updated.LastName == "" && updated.Password == "" {
		c.JSON(http.StatusNoContent, nil) // 204 expects no body in the response
		return
	}

	// Update the user successfully
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
		logger.Error().Msg("Failed to update user")
    return
  }

  c.JSON(http.StatusOK, user)
	logger.Info().Msg("User updated successfully")
}