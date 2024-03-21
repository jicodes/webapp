package initializers

import (
	"bufio"
	"os"
	"strings"

	"github.com/jicodes/webapp/internals/logger"
	"github.com/joho/godotenv"
)

var initializersLogger = logger.GetLogger().With().Str("service", "initializers").Logger()

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
			initializersLogger.Error().Err(err).Msg("Error loading .env file")
		}
	}
}

func LoadAppProperties() {
	file, err := os.Open("/opt/myapp/app.properties")
	if err != nil {
		initializersLogger.Error().Err(err).Msg("Error opening app.properties file")
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
		initializersLogger.Error().Err(err).Msg("Error reading app.properties file")
	}
}