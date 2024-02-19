package main

import (
	"bytes"
	"fmt"

	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

const (
	testFirstName = "First"
	testLastName = "Last"
	testPassword = "TestPassword"
	testUsername = "testuser@example.com"
)

const (
	updatedFirstName = "UpdatedFirst"
	updatedLastName = "UpdatedLast"
	updatedPassword = "UpdatedPassword"
)

func setupUser(router *gin.Engine, firstName, lastName, password, username string) (int, error) {
	// Create an user 
	createUserPayload := []byte(fmt.Sprintf(`{"first_name" : "%s", "last_name" : "%s", "password" : "%s", "username": "%s"}`, firstName, lastName, password, username))
	createUserReq, err := http.NewRequest("POST", "/v1/user", bytes.NewBuffer(createUserPayload))
	if err != nil {
			return 0, fmt.Errorf("Failed to create request: %v", err)
	}
	createUserReq.Header.Set("Content-Type", "application/json")

	createUserResp := httptest.NewRecorder()

	router.ServeHTTP(createUserResp, createUserReq)

	if createUserResp.Code != http.StatusCreated {
			return createUserResp.Code, fmt.Errorf("Failed to create user: status code %v", createUserResp.Code)
	}

	return createUserResp.Code, nil
}

func cleanupUser(username string) error {
	user := models.User{}
	err := initializers.DB.First(&user, "username = ?", username).Error
	if err != nil {
			return fmt.Errorf("Failed to find user: %v", err)
	}

	err = initializers.DB.Delete(&user).Error
	if err != nil {
			return fmt.Errorf("Failed to delete user: %v", err)
	}

	return nil
}


func TestCreateUser(t *testing.T) {
	// Initialize your Gin router
	router := setupRouter()

	// Create an user 
	resCode, createUserErr := setupUser(router, testFirstName, testLastName, testPassword, testUsername)
	if createUserErr != nil {
			t.Fatalf("Failed to create user: %v", createUserErr)
	}

	assert.Equal(t, http.StatusCreated, resCode, "Should return a 201 status code for a successful user creation")
	// Validate the user was created 
	getUserReq, _ := http.NewRequest("GET", "/v1/user/self", nil)
	getUserReq.Header.Set("Content-Type", "application/json")

	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(testUsername + ":" + testPassword))
	getUserReq.Header.Set("Authorization", basicAuth)

	getUserResp := httptest.NewRecorder()
	router.ServeHTTP(getUserResp, getUserReq)

	assert.Equal(t, http.StatusOK, getUserResp.Code)

	var createdUser models.User
	err := json.Unmarshal(getUserResp.Body.Bytes(), &createdUser)
	if err != nil {
    t.Fatalf("Failed to parse response body: %v", err)
	}

	expectedUser := models.User{
		FirstName: testFirstName,
		LastName:  testLastName,
		Username:  testUsername,
	}

	assert.Equal(t, expectedUser.LastName, createdUser.FirstName)
	assert.Equal(t, expectedUser.LastName, createdUser.LastName)
	assert.Equal(t, expectedUser.Username, createdUser.Username)

	cleanupUser(testUsername)
}

func TestUpdateUser(t *testing.T) {
	router := setupRouter()

	resCode, _ := setupUser(router, testFirstName, testLastName, testPassword, testUsername)
	assert.Equal(t, http.StatusCreated, resCode, "Should return a 201 status code for a successful user creation")

	// Test 2 - Update the account and validate the account was updated
	updateUserPayload := []byte(fmt.Sprintf(`{"first_name" : "%s", "last_name" : "%s", "password" : "%s"}`, updatedFirstName, updatedLastName, updatedPassword))
	updateUserReq, _ := http.NewRequest("PUT", "/v1/user/self", bytes.NewBuffer(updateUserPayload))
	updateUserReq.Header.Set("Content-Type", "application/json")

	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(testUsername + ":" + testPassword))
	updateUserReq.Header.Set("Authorization", basicAuth)

	updateUserResp := httptest.NewRecorder()
	router.ServeHTTP(updateUserResp, updateUserReq)

	assert.Equal(t, http.StatusOK, updateUserResp.Code)

	// Validate the updated user information
	var updatedUser models.User
    err := json.Unmarshal(updateUserResp.Body.Bytes(), &updatedUser)
    if err != nil {
        t.Fatalf("Failed to parse response body: %v", err)
    }

	assert.Equal(t, updatedFirstName, updatedUser.FirstName)
	assert.Equal(t, updatedLastName, updatedUser.LastName)

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(updatedPassword))
	assert.NoError(t, comparePasswordErr)
}
