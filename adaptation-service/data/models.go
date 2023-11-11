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
		Instructions:      Instructions{},
		UsersInstructions: UsersInstructions{},
	}
}

type Models struct {
	Instructions      Instructions
	UsersInstructions UsersInstructions
}

type Instructions struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Path        string    `json:"path"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UsersInstructions struct {
	UserID        int       `json:"user_id"`
	InstructionID int       `json:"instruction_id"`
	SolvedAt      time.Time `json:"solved_at,omitempty"`
}

type SolvedInstructions struct {
	Instructions
	Solved bool `json:"solved"`
}

func (i *Instructions) GetAll() ([]*Instructions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, description, path, created_at, updated_at from instructions order by id`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var instructions []*Instructions

	for rows.Next() {
		var instruction Instructions
		err = rows.Scan(
			&instruction.ID,
			&instruction.Title,
			&instruction.Description,
			&instruction.Path,
			&instruction.CreatedAt,
			&instruction.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		instructions = append(instructions, &instruction)
	}

	return instructions, nil
}

func (i *Instructions) Insert(instruction Instructions) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var newID int
	stmt := `insert into instructions (title, description, path, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id`

	err := db.QueryRowContext(ctx, stmt,
		instruction.Title,
		instruction.Description,
		instruction.Path,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (i *Instructions) GetOne(id int) (*Instructions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, description, path, created_at, updated_at from instructions where id = $1`

	var instructions Instructions
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&instructions.ID,
		&instructions.Title,
		&instructions.Description,
		&instructions.Path,
		&instructions.CreatedAt,
		&instructions.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &instructions, nil
}

func (i *Instructions) Update(userID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `update users_instructions set solved_at = $1
		where user_id = $2 and instruction_id = $3`

	_, err := db.ExecContext(ctx, stmt,
		time.Now(),
		userID,
		i.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (ui *UsersInstructions) Insert(usersKnowledges UsersInstructions) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `insert into users_instructions (user_id, instruction_id) values ($1, $2) returning user_id`

	var userID int
	err := db.QueryRowContext(ctx, stmt,
		usersKnowledges.UserID,
		usersKnowledges.InstructionID,
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (ui *UsersInstructions) GetAll(id int) ([]*Instructions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select instructions.id, instructions.title, instructions.description, instructions.path,
       	instructions.created_at, instructions.updated_at
		from instructions, users_instructions
		where users_instructions.instruction_id = instructions.id
		and users_instructions.user_id = $1;`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var instructions []*Instructions

	for rows.Next() {
		var instruction Instructions
		err = rows.Scan(
			&instruction.ID,
			&instruction.Title,
			&instruction.Description,
			&instruction.Path,
			&instruction.CreatedAt,
			&instruction.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		instructions = append(instructions, &instruction)
	}

	return instructions, nil
}

func (ui *UsersInstructions) GetAllSolved(id int) ([]*int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select instruction_id from users_instructions where user_id = $1 and solved_at is not null`

	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersInstructionIDs []*int

	for rows.Next() {
		var instructionID int
		err = rows.Scan(
			&instructionID,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		usersInstructionIDs = append(usersInstructionIDs, &instructionID)
	}

	return usersInstructionIDs, nil
}

func (ui *UsersInstructions) GetPercent(id int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	//TODO
	var countUI int
	query := `select count(*) from users_instructions where user_id = $1 and solved_at is not null`
	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&countUI,
	)
	if err != nil {
		return 0, err
	}

	var countI int
	query = `select count(*) from users_instructions where user_id = $1`
	row = db.QueryRowContext(ctx, query, id)

	err = row.Scan(
		&countI,
	)
	if err != nil {
		return 0, err
	}

	percent := countUI * 100 / countI

	return percent, nil
}
