package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{}
}

type Models struct {
	Info           Knowledge
	UsersKnowledge UsersKnowledge
}

type Knowledge struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UsersKnowledge struct {
	UserID      int       `json:"user_id"`
	KnowledgeID int       `json:"knowledge_id"`
	SolvedAt    time.Time `json:"solved_at,omitempty"`
}
