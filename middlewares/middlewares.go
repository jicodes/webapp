package middlewares

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"

	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/models"
)

func CheckRequestMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if logErr != nil {
			panic(logErr)
		}
		defer logFile.Close()
		logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

		if c.Request.URL.Path == "/healthz" && c.Request.Method != http.MethodGet {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			logger.Error().Msg("Method for healthz api not allowed")
			return
		}
		c.Next()
	}
}

func CheckPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if logErr != nil {
			panic(logErr)
		}
		defer logFile.Close()
		logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

		if c.Request.URL.Path == "/healthz" && (c.Request.ContentLength > 0 || len(c.Request.URL.Query()) > 0) {
			c.AbortWithStatus(http.StatusBadRequest)
			logger.Error().Msg("Payload for healthz api not allowed")
			return
		}
		c.Next()
	}
}

func BasicAuth () gin.HandlerFunc {
	return func(c *gin.Context) {
		logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if logErr != nil {
			panic(logErr)
		}
		defer logFile.Close()
		logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Header("WWW-Authenticate", "Basic")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	
    encodedCreds := strings.TrimPrefix(authHeader, "Basic ")
    decoded, err := base64.StdEncoding.DecodeString(encodedCreds)
    if err != nil {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid base64 string",
      })
      return
    }
    creds := string(decoded)
    credentials := strings.Split(creds, ":")
    if len(credentials) != 2 {
			logger.Error().Err(err).Msg("Invalid credentials")
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
        "error": "Invalid credentials",
      })
      return
    }
    username := credentials[0]
    password := credentials[1]

		var user models.User
		result := initializers.DB.First(&user, "username = ?", username)
		// Check if the user exists 
		if result.Error != nil {
			logger.Error().Err(result.Error).Msg("User does not exist")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User does not exist",
			})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			c.Header("WWW-Authenticate", "Basic")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Username or password is incorrect",
			})
			logger.Error().Err(err).Msg("Username or password is incorrect")
			return
		}
		
		c.Set("user", user)

		c.Next()
	}
}