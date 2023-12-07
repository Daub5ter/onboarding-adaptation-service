package main

import (
	"broker/event"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RequestPayload struct {
	Action            string                   `json:"action"`
	Auth              AuthUserPayload          `json:"auth,omitempty"`
	Session           SessionTokenPayload      `json:"session,omitempty"`
	Reg               RegUserPayload           `json:"reg,omitempty"`
	Email             EmailPayload             `json:"email,omitempty"`
	ID                IDPayload                `json:"id,omitempty"`
	Known             KnowledgePayload         `json:"known,omitempty"`
	UsersKnown        UsersKnowledgesPayload   `json:"users_known,omitempty"`
	Instruction       InstructionPayload       `json:"instruction,omitempty"`
	UsersInstructions UsersInstructionsPayload `json:"users_instructions,omitempty"`
	Mail              MailPayload              `json:"mail,omitempty"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type AuthUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SessionTokenPayload stores token of user`s session
type SessionTokenPayload struct {
	SessionToken string `json:"session_token"`
}

type RegUserPayload struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Password   string `json:"password"`
	Profession string `json:"profession"`
}

type EmailPayload struct {
	Email string `json:"email"`
}

type IDPayload struct {
	ID int `json:"id"`
}

type KnowledgePayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type InstructionPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

type UsersKnowledgesPayload struct {
	UserID      int `json:"user_id"`
	KnowledgeID int `json:"knowledge_id"`
}

type UsersInstructionsPayload struct {
	UserID        int `json:"user_id"`
	InstructionID int `json:"instruction_id"`
}

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth_user":
		app.authenticateUserViaRabbit(w, requestPayload.Auth)
	case "authenticate_user_session":
		app.authenticateUserSessionViaRabbit(w, requestPayload.Session)
	case "reg_user":
		app.registrateUser(w, requestPayload.Reg)
	case "get_all_user":
		app.getAllUser(w)
	case "get_user_by_email":
		app.getUserByEmailViaRabbit(w, requestPayload.Email)
	case "get_user_by_id":
		app.getByIDUser(w, requestPayload.ID)
	case "get_all_knowledge":
		app.getAllKnowledge(w, requestPayload.ID)
	case "get_percent_knowledge":
		app.getPercentKnowledge(w, requestPayload.ID)
	case "add_knowledge":
		app.addKnowledge(w, requestPayload.Known)
	case "add_users_knowledge":
		app.addUsersKnowledge(w, requestPayload.UsersKnown)
	case "get_instruction":
		app.getInstruction(w, requestPayload.ID)
	case "get_all_instructions":
		app.getAllInstructions(w)
	case "get_users_instructions":
		app.getUsersInstructions(w, requestPayload.ID)
	case "add_instruction":
		app.addInstruction(w, requestPayload.Instruction)
	case "add_users_instruction":
		app.addUsersInstruction(w, requestPayload.UsersInstructions)
	case "solve_instruction":
		app.solveInstruction(w, requestPayload.UsersInstructions)
	case "get_percent_instructions":
		app.getPercentInstructions(w, requestPayload.ID)
	case "mail":
		app.sendMail(w, requestPayload.Mail)

	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

// authenticateUserSession auths user if session is valid
func (app *Config) authenticateUserSession(w http.ResponseWriter, st SessionTokenPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(st, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate_session", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusOK)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Session is valid"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// authenticateUser auth user with email and password
func (app *Config) authenticateUser(w http.ResponseWriter, a AuthUserPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payload)
}

// getAllUser return all users
func (app *Config) getAllUser(w http.ResponseWriter) {
	// call the service
	request, err := http.NewRequest("GET", "http://authentication-service/get_all", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received users"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// getByEmailUser return user by email
func (app *Config) getByEmailUser(w http.ResponseWriter, e EmailPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(e, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/get_by_email", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received user"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// getByIDUser return user by ID
func (app *Config) getByIDUser(w http.ResponseWriter, i IDPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(i, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/get_by_id", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received user"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// registrateUser registrate user and return user`s ID
func (app *Config) registrateUser(w http.ResponseWriter, r RegUserPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(r, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/registrate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusCreated {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Registrated"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusCreated, payload)
}

// getAllKnowledge return user`s solved and unsolved knowledge
func (app *Config) getAllKnowledge(w http.ResponseWriter, i IDPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(i, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://onboarding-service/get_all", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received knowledge"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// getPercentKnowledge return percent of solved knowledge
func (app *Config) getPercentKnowledge(w http.ResponseWriter, i IDPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(i, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://onboarding-service/get_percent", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received percent"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// addKnowledge add the new knowledge
func (app *Config) addKnowledge(w http.ResponseWriter, k KnowledgePayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(k, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://onboarding-service/add_known", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusCreated {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Knowledge added"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusCreated, payload)
}

// addUsersKnowledge add the new users solved knowledge
func (app *Config) addUsersKnowledge(w http.ResponseWriter, uk UsersKnowledgesPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(uk, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://onboarding-service/add_users_known", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusCreated {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Users knowledge added"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusCreated, payload)
}

// getAllInstructions return all instructions
func (app *Config) getAllInstructions(w http.ResponseWriter) {
	// call the service
	request, err := http.NewRequest("GET", "http://adaptation-service/get_all", nil)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received all instructions"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// getUsersInstructions return all users instructions
func (app *Config) getUsersInstructions(w http.ResponseWriter, i IDPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(i, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://adaptation-service/get_users_instructions", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received users instructions"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// addInstruction add the new instruction
func (app *Config) addInstruction(w http.ResponseWriter, i InstructionPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(i, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://adaptation-service/add_instruction", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusCreated {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Instruction added"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusCreated, payload)
}

// addUsersInstruction add the new users instruction
func (app *Config) addUsersInstruction(w http.ResponseWriter, ui UsersInstructionsPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(ui, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://adaptation-service/add_users_instruction", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusCreated {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Users instruction added"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusCreated, payload)
}

// getInstruction return one instruction by ID
func (app *Config) getInstruction(w http.ResponseWriter, i IDPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(i, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://adaptation-service/get_instruction", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received instruction"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// solveInstruction add solved point into db
func (app *Config) solveInstruction(w http.ResponseWriter, ui UsersInstructionsPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(ui, "", "\t")

	// call the service
	request, err := http.NewRequest("PUT", "http://adaptation-service/solve_instruction", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Solved"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// getPercentInstructions return percent of solved knowledge
func (app *Config) getPercentInstructions(w http.ResponseWriter, i IDPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(i, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://adaptation-service/get_percent", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusOK {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into
	var jsonFromService jsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Received percent"
	payload.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusOK, payload)
}

// sendMail send some text to users email
func (app *Config) sendMail(w http.ResponseWriter, msg MailPayload) {
	jsonData, _ := json.MarshalIndent(msg, "", "\t")

	// call the mail service
	mailServiceURL := "http://mailer-service/send"

	// post to mail service
	request, err := http.NewRequest("POST", mailServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the right status code
	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling mail service"))
		return
	}

	// send back json
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Message sent to " + msg.To

	app.writeJSON(w, http.StatusAccepted, payload)
}

// authenticateUserViaRabbit auths user with email and password via RabbitMQ
func (app *Config) authenticateUserViaRabbit(w http.ResponseWriter, a AuthUserPayload) {
	var requestPayload RequestPayload

	requestPayload.Action = "auth_user"
	requestPayload.Auth = a

	response, err := app.pushToQueue(requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse

	err = json.Unmarshal(response, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// authenticateUserSessionViaRabbit auths user if session is valid via RabbitMQ
func (app *Config) authenticateUserSessionViaRabbit(w http.ResponseWriter, st SessionTokenPayload) {
	var requestPayload RequestPayload

	requestPayload.Action = "authenticate_user_session"
	requestPayload.Session = st

	response, err := app.pushToQueue(requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse

	err = json.Unmarshal(response, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, payload)
}

// getUserByEmailViaRabbit returns user by email via RabbitMQ
func (app *Config) getUserByEmailViaRabbit(w http.ResponseWriter, e EmailPayload) {
	var requestPayload RequestPayload

	requestPayload.Action = "get_user_by_email"
	requestPayload.Email = e

	response, err := app.pushToQueue(requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse

	err = json.Unmarshal(response, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) getAllKnowledgeViaRabbit(w http.ResponseWriter, i IDPayload) {}

// pushToQueue pushes request to queue of RabbitMQ
func (app *Config) pushToQueue(payload RequestPayload) ([]byte, error) {
	var response []byte

	emitter, err := event.NewEventEmitter(payload.Action, app.Rabbit)
	if err != nil {
		return nil, err
	}

	var j []byte

	j, _ = json.MarshalIndent(&payload, "", "\t")

	switch payload.Action {
	case "auth_user":
		response, err = emitter.PushWithResponse(string(j), payload.Action, "auth.user")
	case "get_user_by_email":
		response, err = emitter.PushWithResponse(string(j), payload.Action, "get.user.by.email")
	case "authenticate_user_session":
		response, err = emitter.PushWithResponse(string(j), payload.Action, "authenticate.user.session")
	case "get_all_knowledge":
		response, err = emitter.PushWithResponse(string(j), payload.Action, "get.all.knowledge")
	case "add_users_knowledge":
		response, err = emitter.PushWithResponse(string(j), payload.Action, "add.users.knowledge")
	case "get_percent_knowledge":
		response, err = emitter.PushWithResponse(string(j), payload.Action, "get.percent.knowledge")
	default:
		log.Printf("invalid name of channel RabbitMQ %s", payload.Action)
	}

	if err != nil {
		return nil, err
	}

	return response, err
}
