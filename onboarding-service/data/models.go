package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Knowledge:      Knowledge{},
		UsersKnowledge: UsersKnowledge{},
	}
}

type Models struct {
	Knowledge      Knowledge
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

func (u *Knowledge) GetAll() ([]*Knowledge, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, description, created_at, updated_at from knowledges order by title`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var knowledges []*Knowledge

	for rows.Next() {
		var knowledge Knowledge
		err := rows.Scan(
			&knowledge.ID,
			&knowledge.Title,
			&knowledge.Description,
			&knowledge.CreatedAt,
			&knowledge.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		knowledges = append(knowledges, &knowledge)
	}

	return knowledges, nil
}
