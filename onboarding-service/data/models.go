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
		Knowledge:       Knowledge{},
		UsersKnowledges: UsersKnowledges{},
	}
}

type Models struct {
	Knowledge       Knowledge
	UsersKnowledges UsersKnowledges
}

type Knowledge struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UsersKnowledges struct {
	UserID      int       `json:"user_id"`
	KnowledgeID int       `json:"knowledge_id"`
	SolvedAt    time.Time `json:"solved_at,omitempty"`
}

type SolvedKnowledges struct {
	Knowledge Knowledge `json:"knowledge"`
	Solved    bool      `json:"solved"`
}

type KnowledgeJoinUsersKnowledges struct {
	Knowledge
	UsersKnowledges
}

func (k *Knowledge) GetAll() ([]*Knowledge, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, description, created_at, updated_at from knowledges order by id`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var knowledges []*Knowledge

	for rows.Next() {
		var knowledge Knowledge
		err = rows.Scan(
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

func (k *Knowledge) Insert(knowledge Knowledge) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into knowledges (title, description, created_at, updated_at) values ($1, $2, $3, $4) returning id`

	err := db.QueryRowContext(ctx, stmt,
		knowledge.Title,
		knowledge.Description,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (k *Knowledge) GetOne(id int) (*Knowledge, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, description, created_at, updated_at from knowledges where id = $1`

	var knowledge Knowledge
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&knowledge.ID,
		&knowledge.Title,
		&knowledge.Description,
		&knowledge.CreatedAt,
		&knowledge.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &knowledge, nil
}

func (uk *UsersKnowledges) Insert(usersKnowledges UsersKnowledges) (time.Time, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into users_knowledges (user_id, knowledge_id, solved_at) values ($1, $2, $3) returning solved_at`

	var solvedAt time.Time
	err := db.QueryRowContext(ctx, stmt,
		usersKnowledges.UserID,
		usersKnowledges.KnowledgeID,
		time.Now(),
	).Scan(&solvedAt)

	if err != nil {
		return time.Time{}, err
	}

	return solvedAt, nil
}

func (uk *UsersKnowledges) GetPercent(id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//TODO
	var countUK int
	query := `select count(*) from users_knowledges where user_id = $1`
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&countUK,
	)
	if err != nil {
		log.Println("wtf", err)
		return 0, err
	}

	var countK int
	query = `select count(*) from knowledges`
	row = db.QueryRowContext(ctx, query)

	err = row.Scan(
		&countK,
	)
	if err != nil {
		return 0, err
	}

	percent := countUK * 100 / countK

	return percent, nil
}

func (uk *UsersKnowledges) GetAll(id int) ([]*int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select knowledge_id from users_knowledges where user_id = $1`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersKnowledgeID []*int

	for rows.Next() {
		var knowledgeID int
		err = rows.Scan(
			&knowledgeID,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		usersKnowledgeID = append(usersKnowledgeID, &knowledgeID)
	}

	return usersKnowledgeID, nil
}
