package main

import (
	"adaptation/data"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// GetAll return all instructions
func (app *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	instructions, err := app.Models.Instructions.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log GetAll
	go app.logRequest("get all instructions", fmt.Sprintf("got %v instructions", len(instructions)))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Received %v instructions", len(instructions)),
		Data:    instructions,
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
		resp.Instructions = instruction
		resp.Solved = solved

		response = append(response, &resp)
	}

	// log GetUsersInstructions
	go app.logRequest("get user instructions", fmt.Sprintf("got %v instructions", len(response)))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received all instructions"),
		Data:    response,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetInstructionByID return instructions from db
func (app *Config) GetInstructionByID(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	instruction, err := app.Models.Instructions.GetOne(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log GetInstructionByID
	go app.logRequest("get instruction by id", fmt.Sprintf("got %s", instruction.Title))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received instruction"),
		Data:    instruction,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// AddInstruction insert instruction into db
func (app *Config) AddInstruction(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.Instructions

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	id, err := app.Models.Instructions.Insert(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log AddInstruction
	go app.logRequest("add instruction", fmt.Sprintf("%s added", requestPayload.Title))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created instruction with id %v", id),
		Data:    id,
	}

	app.writeJSON(w, http.StatusCreated, payload)
}

// AddUsersInstruction insert solved instruction into db
func (app *Config) AddUsersInstruction(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.UsersInstructions

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userID, err := app.Models.UsersInstructions.Insert(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log AddUsersInstruction
	go app.logRequest("add users instruction", fmt.Sprintf("added users with id %v instruction with id %v", userID, requestPayload.InstructionID))

	payload := JsonResponse{
		Error: false,
		Message: fmt.Sprintf("Created users instruction with user id %v, instruction id %v",
			requestPayload.UserID, requestPayload.InstructionID),
		Data: userID,
	}

	app.writeJSON(w, http.StatusCreated, payload)
}

// SolveInstruction solve users instruction into db
func (app *Config) SolveInstruction(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.UsersInstructions

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	instruction, err := app.Models.Instructions.GetOne(requestPayload.InstructionID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	err = instruction.Update(requestPayload.UserID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// log SolveInstruction
	go app.logRequest("solve instruction", fmt.Sprintf("solved instruction %s", instruction.Title))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Solved instruction with id %v", instruction.ID),
		Data:    requestPayload.UserID,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) GetPercent(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	percent, err := app.Models.UsersInstructions.GetPercent(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log GetPercent
	go app.logRequest("get percent instruction", fmt.Sprintf("got percent %v", percent))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received percent of instructions with user`s id: %v", requestPayload.ID),
		Data:    percent,
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
