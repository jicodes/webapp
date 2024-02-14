package middlewares

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/models"
	"golang.org/x/crypto/bcrypt"
)

func CheckRequestMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/healthz" && c.Request.Method != http.MethodGet {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		c.Next()
	}
}

func CheckPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/healthz" && (c.Request.ContentLength > 0 || len(c.Request.URL.Query()) > 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Next()
	}
}

func BasicAuth () gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Header("WWW-Authenticate", "Basic")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	
    encodedCreds := strings.TrimPrefix(authHeader, "Basic ")
    decoded, err := base64.StdEncoding.DecodeString(encodedCreds)
    if err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid base64 string",
      })
      c.Abort()
      return
    }
    creds := string(decoded)
    credentials := strings.Split(creds, ":")
    if len(credentials) != 2 {
      c.JSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid credentials",
      })
      c.Abort()
      return
    }
    username := credentials[0]
    password := credentials[1]

		var user models.User
		result := initializers.DB.First(&user, "username = ?", username)
		// Check if the user exists 
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User does not exist",
			})
			c.Abort()
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.Header("WWW-Authenticate", "Basic")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			c.Abort()
			return
		}
		
		c.Set("user", user)

		c.Next()
	}
}