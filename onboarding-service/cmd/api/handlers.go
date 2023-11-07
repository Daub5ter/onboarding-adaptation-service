package main

import (
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

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received percent of knowledge with user`s id: %v", requestPayload.ID),
		Data:    percent,
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// GetAll return user`s solved and unsolved knowledge
func (app *Config) GetAll(w http.ResponseWriter, r *http.Request) {
	knowledges, err := app.Models.Knowledge.GetAll()
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
	}

	//knownPayload, _ := json.Marshal(knowleages)
	//IDsPayload, _ := json.Marshal(solvedKnowleages)

	//bytes.NewBuffer(knownPayload)
	//bytes.NewBuffer(IDsPayload)

	log.Println("log.Println(&knowledges)")
	log.Println(&knowledges)
	log.Println(knowledges)
	//log.Println(IDsPayload)

	type responsePayload struct {
		knowledge data.Knowledge
		solved    bool
	}

	var response []*responsePayload

	/*for _, knowledge := range knownPayload
		var solved bool

		for _, ukID := range IDsPayload {
			if knowledge == ukID {
				solved = true
				break
			}
		}

		var resp responsePayload
		resp.knowledge = *knowledge
		resp.solved = solved

		response = append(response, &resp)
	}*/

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
	}

	usersKnowledge, err := app.Models.UsersKnowledges.GetAll(requestPayload.ID)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
	}

	payload := JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("received all knowledge"),
		Data:    usersKnowledge,
	}

	app.writeJSON(w, http.StatusOK, payload)
}
