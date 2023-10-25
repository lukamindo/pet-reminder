package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/response"
)

type (
	Pet struct {
		ID          int        `db:"id"`
		Name        string     `db:"name"`
		OwnerID     int        `db:"owner_id"`
		DateOfBirth *time.Time `db:"date_of_birth"`
		CreatedAt   time.Time  `db:"created_at"`
	}

	Pets []Pet
)

// PetCreate creates Pet
func PetCreate(c context.Context, db sqlx.ExtContext, p Pet) (*int, error) {
	query, args, err := sqlx.Named(`
	INSERT INTO pets
		(name
		,owner_id
		,date_of_birth
		)
	VALUES
		(:name
		,:owner_id
		,:date_of_birth)
	RETURNING id`, p)
	if err != nil {
		return nil, err
	}
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	var petID int
	err = db.QueryRowxContext(c, query, args...).Scan(&petID)
	if err != nil {
		return nil, err
	}
	return &petID, nil
}

// PetList returns all pets
func PetList(c context.Context, db sqlx.ExtContext) (*Pets, error) {
	var ps Pets
	err := sqlx.SelectContext(c, db, &ps, `
		SELECT 
			id,
			name,
			owner_id,
			date_of_birth,
			created_at		
		FROM pets`)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &ps, nil
}

// PetByID returns pet by ID
func PetByID(c context.Context, db sqlx.ExtContext, id int) (*Pet, error) {
	var p Pet
	err := sqlx.GetContext(c, db, &p, `
		SELECT 
			id,
			name,
			owner_id,
			date_of_birth,
			created_at		
		FROM pets
		WHERE id = $1`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// PetUpdate updates pet
func PetUpdate(c context.Context, db sqlx.ExtContext, p Pet) error {
	_, err := sqlx.NamedExecContext(c, db, `
	UPDATE pets
	SET
		name = :name,
		date_of_birth = :date_of_birth
	WHERE id = :id`, p)
	if err != nil {
		return err
	}
	return nil
}

// PetDeleteByID removes pet by id
func PetDeleteByID(c context.Context, db sqlx.ExtContext, id int) (*int64, error) {
	result, err := db.ExecContext(c, `
		DELETE FROM pets
		WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, _ := result.RowsAffected()
	return &rowsAffected, err
}

// Response converts Pet to response.Pet
func (p Pet) Response() response.Pet {
	return response.Pet{
		ID:          p.ID,
		Name:        p.Name,
		OwnerID:     p.OwnerID,
		DateOfBirth: p.DateOfBirth,
		CreatedAt:   p.CreatedAt,
	}
}

// Response converts Pets to response.Pets with pointers
func (ps Pets) Response() *response.Pets {
	ret := make(response.Pets, len(ps))
	for i, pet := range ps {
		ret[i] = pet.Response()
	}
	return &ret
}
