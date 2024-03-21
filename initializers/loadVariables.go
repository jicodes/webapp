package initializers

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
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
	file, err := os.Open("/opt/myapp/app.properties")
	if err != nil {
		log.Fatal("Error opening app.properties file")
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
		log.Fatal("Error reading app.properties file")
	}
}