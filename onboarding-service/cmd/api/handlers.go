package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"onboarding/data"
)

// GetKnownByID return knowledge from db
func (app *Config) GetKnownByID(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	knowledge, err := app.Models.Knowledge.GetOne(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log GetKnownByID
	go app.logRequest("get knowledge by id", fmt.Sprintf("got knowledge %s", knowledge.Title))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received knowledge"),
		Data:    knowledge,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// AddKnown insert knowledge into db
func (app *Config) AddKnown(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.Knowledge

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	id, err := app.Models.Knowledge.Insert(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log AddKnown
	go app.logRequest("add knowledge", fmt.Sprintf("got knowledge with id %v", id))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created knowledge with id %s", id),
		Data:    id,
	}

	app.writeJSON(w, http.StatusCreated, payload)
}

// AddUsersKnown insert solved knowledge into db
func (app *Config) AddUsersKnown(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.UsersKnowledges

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	solvedAt, err := app.Models.UsersKnowledges.Insert(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log AddUsersKnown
	go app.logRequest("add users knowledge", fmt.Sprintf("user with id %v added knowledge with id %v", requestPayload.UserID, requestPayload.KnowledgeID))

	payload := JsonResponse{
		Error: false,
		Message: fmt.Sprintf("Created users knowledge with user id %v, knowledge id %v",
			requestPayload.UserID, requestPayload.KnowledgeID),
		Data: solvedAt,
	}

	app.writeJSON(w, http.StatusCreated, payload)
}

// GetPercentKnown return percent of solved knowledge
func (app *Config) GetPercentKnown(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	percent, err := app.Models.UsersKnowledges.GetPercent(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log GetPercentKnown
	go app.logRequest("get percent knowledge", fmt.Sprintf("got percent %v", percent))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received percent of knowledge with user`s id: %v", requestPayload.ID),
		Data:    percent,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetAll return user`s solved and unsolved knowledge
func (app *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	var knowns []data.Knowledge
	var ids []int

	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	solvedKnowleages, err := app.Models.UsersKnowledges.GetAll(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	knowledges, err := app.Models.Knowledge.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	knownPayload, _ := json.Marshal(knowledges)
	IDsPayload, _ := json.Marshal(solvedKnowleages)

	bytes.NewBuffer(knownPayload)
	bytes.NewBuffer(IDsPayload)

	err = json.Unmarshal(knownPayload, &knowns)
	if err != nil {
		app.errorJSON(w, errors.New("error with unmarshal"), http.StatusBadGateway)
		return
	}
	err = json.Unmarshal(IDsPayload, &ids)
	if err != nil {
		app.errorJSON(w, errors.New("error with unmarshal"), http.StatusBadGateway)
		return
	}

	var response []*data.SolvedKnowledges

	for _, knowledge := range knowns {
		var solved bool

		for _, ukID := range ids {
			if knowledge.ID == ukID {
				solved = true
				break
			}
		}

		var resp data.SolvedKnowledges
		resp.Knowledge = knowledge
		resp.Solved = solved

		response = append(response, &resp)
	}

	// log GetPercentKnown
	go app.logRequest("get all knowledge", fmt.Sprintf("got %v knowledges", len(response)))

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received all knowledge"),
		Data:    response,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetAllKnown unused func
func (app *Config) GetAllKnown(w http.ResponseWriter, r *http.Request) {
	knowledges, err := app.Models.Knowledge.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received all knowledge"),
		Data:    knowledges,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetAllUsersKnown unused func
func (app *Config) GetAllUsersKnown(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		ID int `json:"id"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	usersKnowledge, err := app.Models.UsersKnowledges.GetAll(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received all knowledge"),
		Data:    usersKnowledge,
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
