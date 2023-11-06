package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := app.Models.Knowledge.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Received %v knowledges", len(users)),
		Data:    users,
	}

	app.writeJSON(w, http.StatusOK, payload)
}
