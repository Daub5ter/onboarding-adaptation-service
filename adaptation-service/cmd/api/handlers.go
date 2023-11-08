package main

import (
	"adaptation/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetAll return all instructions
func (app *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := app.Models.Instructions.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Received %v instructions", len(users)),
		Data:    users,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetUsersInstructions return user`s solved and unsolved instructions
func (app *Config) GetUsersInstructions(w http.ResponseWriter, r *http.Request) {
	var instrs []data.Instructions
	var ids []int

	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	solvedInstructions, err := app.Models.UsersInstructions.GetAllSolved(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	instructions, err := app.Models.UsersInstructions.GetAll(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	instrsPayload, _ := json.Marshal(instructions)
	IDsPayload, _ := json.Marshal(solvedInstructions)

	bytes.NewBuffer(instrsPayload)
	bytes.NewBuffer(IDsPayload)

	err = json.Unmarshal(instrsPayload, &instrs)
	if err != nil {
		app.errorJSON(w, errors.New("error with unmarshal"), http.StatusBadGateway)
		return
	}
	err = json.Unmarshal(IDsPayload, &ids)
	if err != nil {
		app.errorJSON(w, errors.New("error with unmarshal"), http.StatusBadGateway)
		return
	}

	var response []*data.SolvedInstructions

	for _, instruction := range instrs {
		var solved bool

		for _, uiID := range ids {
			if instruction.ID == uiID {
				solved = true
				break
			}
		}

		var resp data.SolvedInstructions
		resp.Instruction = instruction
		resp.Solved = solved

		response = append(response, &resp)
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received all knowledge"),
		Data:    response,
	}

	app.writeJSON(w, http.StatusOK, payload)
}
