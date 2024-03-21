# Webapp inplemented in Go/Gin with Postgres and Gorm

## Prerequisites

- Go 1.21.6
- Postgres 16

## Packages Installation

````sh
# for web server and routing - Gin
go get -u github.com/gin-gonic/gin

# for loading environment variables
go get -u github.com/joho/godotenv

# for DB
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# for hashing the password
go get -u golang.org/x/crypto/bcrypt

# for testing
go get github.com/stretchr/testify

# for logging
go get -u github.com/rs/zerolog/log

## Note

1. We need use cross build for Go app to build then run on a different environment.

```sh
# for local dev: build on local machine(MacOS) and run on VM server (Linux-CentOS)
GOOS=linux GOARCH=amd64 go build -o webapp

# for workflow: build in Github actions runner (latest-Ubuntu) and run on VM server (Linux-CentOS)
# use static linking to create a standalone binary that doesn't depend on dynamic libraries.
go build -ldflags="-linkmode external -extldflags -static" -o webapp
````
