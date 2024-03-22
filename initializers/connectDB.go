package initializers

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func ConnectDB() () {

	logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if logErr != nil {
		panic(logErr)
	}
	defer logFile.Close()
	logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error().Err(err).Msg("Error connecting to database")
	} else {
		logger.Info().Msg("Connected to database")
	}
}