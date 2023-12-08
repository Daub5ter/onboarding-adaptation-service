package main

import (
	"authentication/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

// GetByEmail return user by email
func (app *Config) GetByEmail(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email string `json:"email"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// get user form database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("recived user"),
		Data:    user,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetByID return user by ID
func (app *Config) GetByID(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetOne(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received user"),
		Data:    user,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetAll return all users from db
func (app *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	// get all users from database
	users, err := app.Models.User.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log GetAll
	go app.logRequest("get all users", fmt.Sprintf("got %v users", len(users)))

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Received %v users", len(users)),
		Data:    users,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// Registrate input user into db
func (app *Config) Registrate(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.User

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// insert user in database
	id, err := app.Models.User.Insert(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log Registrate
	go app.logRequest("registration user", fmt.Sprintf("%s registreted", requestPayload.Email))

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created user with id %s", id),
		Data:    id,
	}

	app.writeJSON(w, http.StatusCreated, payload)
}

// Authenticate auth user with email and password
func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmailWithPassword(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	jwtToken, err := app.Models.UserJWT.CreateJWTToken(user.Email, user.ID)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	u := struct {
		ID           int       `json:"id"`
		Email        string    `json:"email"`
		FirstName    string    `json:"first_name,omitempty"`
		LastName     string    `json:"last_name,omitempty"`
		Profession   string    `json:"profession"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		SessionToken string    `json:"session_token"`
	}{}

	u.ID = user.ID
	u.Email = user.Email
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Profession = user.Profession
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	u.SessionToken = jwtToken

	// log Authenticate
	go app.logRequest("authenticate user", fmt.Sprintf("%s authenticated", requestPayload.Email))

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    u,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// AuthenticateSession checks valid or not session of user
func (app *Config) AuthenticateSession(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		SessionToken string `json:"session_token"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.UserJWT.CheckJWTToken(requestPayload.SessionToken)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// log AuthenticateSession
	go app.logRequest("authenticate user session", fmt.Sprintf("%s authenticated", user.Email))

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("session is valid"),
		Data:    user,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// logRequest requests of logger-service to log event
func (app *Config) logRequest(name, data string) {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
	}
}
