package request

import (
	"time"

	"github.com/lukamindo/pet-reminder/app/db"
)

type PetCreate struct {
	Name        string     `json:"name" validate:"required"`
	OwnerID     int        `json:"owner_id" validate:"required"`
	DateOfBirth *time.Time `json:"date_of_birth" ` //TODO: add time validation on DateOfBirth
}

type PetUpdate struct {
	Name        string     `json:"name"`
	DateOfBirth *time.Time `json:"date_of_birth" `
}

// DB converts request to db object
func (pc PetCreate) DB() db.Pet {
	return db.Pet{
		Name:        pc.Name,
		OwnerID:     pc.OwnerID,
		DateOfBirth: pc.DateOfBirth,
	}
}

func (pc PetUpdate) DB() db.Pet {
	return db.Pet{
		Name:        pc.Name,
		DateOfBirth: pc.DateOfBirth,
	}
}
