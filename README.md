# Webapp inplemented in Go/Gin with Postgres and Gorm

## Prerequisites
- Go 1.21 or higher

## Packages Installation
```sh
# for web server and routing - Gin
go get -u github.com/gin-gonic/gin

# for loading environment variables
go get -u github.com/joho/godotenv

# for DB
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# for hashing the password
go get -u golang.org/x/crypto/bcrypt
```