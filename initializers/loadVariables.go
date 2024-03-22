package initializers

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func LoadVariables() {	

	env := os.Getenv("ENV")

	switch env {
	case "cloud":
		LoadAppProperties()
	case "github":
		break
	default:
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func LoadAppProperties() {
	logFile, err := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

	file, err := os.Open("/opt/myapp/app.properties")
	if err != nil {
		logger.Error().Err(err).Msg("Error opening app.properties file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equalIndex := strings.Index(line, "=")
		if equalIndex == -1 {
			continue
		}
		key := line[:equalIndex]
		value := line[equalIndex+1:]
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		logger.Error().Err(err).Msg("Error reading app.properties file")
	} else {
		logger.Info().Msg("Loaded app.properties file")}
}