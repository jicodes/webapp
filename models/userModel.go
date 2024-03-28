package models

import (
	"regexp"
	"strings"
	"time"
)

type User struct {
    ID                  string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey;->" json:"id"`
    FirstName           string    `gorm:"not null" json:"first_name"`
    LastName            string    `gorm:"not null" json:"last_name"`
    Password            string    `gorm:"not null" json:"password"`
    Username            string    `gorm:"unique;not null" json:"username"`
    Verified            bool      `gorm:"default:false" json:"verified"`
    VerificationToken   string    `gorm:"default:null" json:"verification_token"`
    VerificationTokenCreated time.Time `gorm:"default:null" json:"verification_token_created"`
    AccountCreated time.Time      `gorm:"default:current_timestamp" json:"account_created"`
    AccountUpdated time.Time      `gorm:"default:current_timestamp" json:"account_updated"`
}

type PublicUser struct {
    ID                  string  `json:"id"`
    FirstName           string  `json:"first_name"`
    LastName            string  `json:"last_name"`
    Username            string  `json:"username"`
    Verified            bool    `json:"verified"`
    VerificationToken   string  `json:"verification_token"`
    VerificationTokenCreated time.Time `json:"verification_token_created"`
    AccountCreated time.Time    `json:"account_created"`
    AccountUpdated time.Time    `json:"account_updated"`
}

func ValidateEmail(username string) bool {
    username = strings.ToLower(username)
    var emailRegex = regexp.MustCompile(`^[a-z0-9._\-]+@[a-z0-9.\-]+\.[a-z]{2,7}$`)
    return emailRegex.MatchString(username)
}