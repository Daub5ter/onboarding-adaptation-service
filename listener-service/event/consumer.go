package event

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

// jsonResponse stores json response from services
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Payload is basic structure to indicate action and data`s structure
type Payload struct {
	Action         string                `json:"action"`
	Auth           AuthUserPayload       `json:"auth,omitempty"`
	Session        SessionTokenPayload   `json:"session,omitempty"`
	Reg            RegUserPayload        `json:"reg,omitempty"`
	UpdateUser     UpdateUserPayload     `json:"update_user,omitempty"`
	ChangePassword ChangePasswordPayload `json:"change_password,omitempty"`
	Email          EmailPayload          `json:"email,omitempty"`
	ID             IDPayload             `json:"id,omitempty"`
	Log            LogPayload            `json:"log,omitempty"`
	Mail           MailPayload           `json:"mail,omitempty"`
}

// MailPayload stores data to send mail to user
type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// AuthUserPayload stores data to authenticate user
type AuthUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegUserPayload stores data to registration user
type RegUserPayload struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password"`
	Active    int    `json:"active"`
}

// UpdateUserPayload stores data to update user
type UpdateUserPayload struct {
	Email       string `json:"email"`
	EmailChange string `json:"email_change"`
	FirstName   string `json:"first_name,omitempty,omitempty"`
	LastName    string `json:"last_name,omitempty,omitempty"`
	Active      int    `json:"active,omitempty"`
}

// ChangePasswordPayload stores data to change password
type ChangePasswordPayload struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

// EmailPayload stores data of email
type EmailPayload struct {
	Email string `json:"email"`
}

// IDPayload stores id data
type IDPayload struct {
	ID int `json:"id"`
}

// LogPayload stores log data
type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// SessionTokenPayload stores token of user`s session
type SessionTokenPayload struct {
	SessionToken string `json:"session_token"`
}

func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

func (consumer *Consumer) Listen() error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	if err = ch.QueueBind(q.Name, "auth.user", "auth_user", false, nil); err != nil {
		return err
	}
	if err = ch.QueueBind(q.Name, "get.user.by.email", "get_user_by_email", false, nil); err != nil {
		return err
	}
	if err = ch.QueueBind(q.Name, "authenticate.user.session", "authenticate_user_session", false, nil); err != nil {
		return err
	}
	if err = ch.QueueBind(q.Name, "get.all.knowledge", "get_all_knowledge", false, nil); err != nil {
		return err
	}
	if err = ch.QueueBind(q.Name, "add.users.knowledge", "add_users_knowledge", false, nil); err != nil {
		return err
	}
	if err = ch.QueueBind(q.Name, "get.percent.knowledge", "get_percent_knowledge", false, nil); err != nil {
		return err
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			err = json.Unmarshal(d.Body, &payload)
			if err != nil {
				log.Println(err)
			}

			response := handlePayload(payload)

			jsonResp, err := json.MarshalIndent(response, "", "\t")
			if err != nil {
				log.Println(err)
			}

			err = ch.PublishWithContext(
				context.TODO(),
				"",
				d.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          jsonResp,
				})
			if err != nil {
				log.Println(err)
			}
		}
	}()

	fmt.Printf("Waiting for messages")
	<-forever

	return nil
}

// handlePayload does request and returns response
func handlePayload(payload Payload) jsonResponse {
	response := jsonResponse{}

	switch payload.Action {

	case "auth_user":
		resp, err := auth(payload)
		if err != nil {
			log.Println(err)
		}
		response = resp

	case "get_user_by_email":
		resp, err := getUserByEmail(payload)
		if err != nil {
			log.Println(err)
		}
		response = resp

	case "authenticate_user_session":
		resp, err := authUserSession(payload)
		if err != nil {
			log.Println(err)
		}
		response = resp

	default:
		errString := fmt.Sprintf("invalid name of function %s, RabbitMQ", payload.Action)
		log.Println(errString)
	}

	return response
}

// auth auths user with email and password via RabbitMQ
func auth(entry Payload) (jsonResponse, error) {
	// create some json we'll send to the auth microservice
	jsonData, err := json.MarshalIndent(entry.Auth, "", "\t")
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("error %v", err)}, err
	}

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("error %v", err)}, err
	}

	return handleSync(request, http.StatusOK)
}

// authUserSession auths user if session is valid via RabbitMQ
func authUserSession(entry Payload) (jsonResponse, error) {
	// create some json we'll send to the auth microservice
	jsonData, err := json.MarshalIndent(entry.Session, "", "\t")
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("error %v", err)}, err
	}

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate_session", bytes.NewBuffer(jsonData))
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("error %v", err)}, err
	}

	return handleSync(request, http.StatusOK)
}

// getUserByEmail returns user by email via RabbitMQ
func getUserByEmail(entry Payload) (jsonResponse, error) {
	// create some json we'll send to the auth microservice
	jsonData, err := json.MarshalIndent(entry.Email, "", "\t")
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("error %v", err)}, err
	}

	// call the service
	request, err := http.NewRequest("POST", "http://authentication-service/get_by_email", bytes.NewBuffer(jsonData))
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("error %v", err)}, err
	}

	return handleSync(request, http.StatusOK)
}

// handleAsync is template of async request
func handleAsync(request *http.Request) error {
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode != http.StatusAccepted {
		return errors.New("service don`t work")
	}

	return nil
}

// handleSync is template of sync request. code is http status which should be if function is work
func handleSync(request *http.Request, code int) (jsonResponse, error) {
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("%v", err)}, err
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		err = errors.New("unauthorized")
		return jsonResponse{Error: true, Message: fmt.Sprintf("%v", err)}, err
	} else if response.StatusCode != code {
		err = errors.New("service don`t work")
		return jsonResponse{Error: true, Message: fmt.Sprintf("%v", err)}, err
	}

	jsonService := jsonResponse{}
	err = json.NewDecoder(response.Body).Decode(&jsonService)
	if err != nil {
		return jsonResponse{Error: true, Message: fmt.Sprintf("%v", err)}, err
	}

	return jsonService, nil
}
