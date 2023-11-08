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
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type UsersInstructions struct {
	UserID        int       `json:"user_id"`
	InstructionID int       `json:"instruction_id"`
	SolvedAt      time.Time `json:"solved_at,omitempty"`
}

type SolvedInstructions struct {
	Instruction Instructions `json:"instructions"`
	Solved      bool         `json:"solved"`
}

func (i *Instructions) GetAll() ([]*Instructions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, title, description, created_at, updated_at from instructions order by id`

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

func (ui *UsersInstructions) GetAll(id int) ([]*int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select instructions.id, instructions.title, instructions.description, instructions.created_at, instructions.updated_at
		from instructions instructions, users_instructions users_instructions
		where users_instructions.instruction_id = instructions.id;`

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
