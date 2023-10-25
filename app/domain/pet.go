package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lukamindo/pet-reminder/app/db"
	"github.com/lukamindo/pet-reminder/app/request"
	"github.com/lukamindo/pet-reminder/app/response"
	"github.com/lukamindo/pet-reminder/helper/server"
	"github.com/lukamindo/pet-reminder/helper/validator"
)

type PetService struct {
	connDB *sqlx.DB
}

func NewPetService(connDB *sqlx.DB) PetService {
	return PetService{
		connDB: connDB,
	}
}

func (s PetService) Create(c context.Context, pcr request.PetCreate) (*response.Pet, error) {
	err := validator.ValidateStruct(pcr)
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}
	pcrDB := pcr.DB()
	pID, err := db.PetCreate(c, s.connDB, pcrDB)
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}
	resp := &response.Pet{
		ID:          *pID,
		Name:        pcr.Name,
		OwnerID:     pcr.OwnerID,
		DateOfBirth: pcr.DateOfBirth,
		//TODO: ეს რამდენად უნდა???????? და თუ კი ბაზიდან დავაბრუნებინო ესეც
		CreatedAt: time.Now().UTC(),
	}
	return resp, nil
}

func (s PetService) List(c context.Context) (*response.Pets, error) {
	petList, err := db.PetList(c, s.connDB)
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}
	return petList.Response(), nil
}

func (s PetService) ByID(c context.Context, id int) (*response.Pet, error) {
	pet, err := db.PetByID(c, s.connDB, id)
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}
	if pet == nil {
		return nil, server.ErrBadRequest(fmt.Errorf("bad pet id"))
	}
	resp := pet.Response()
	return &resp, nil
}

func (s PetService) Update(c context.Context, id int, pur request.PetUpdate) (*response.Pet, error) {
	err := validator.ValidateStruct(pur)
	if err != nil {
		return nil, server.ErrBadRequest(err)
	}
	pet, err := db.PetByID(c, s.connDB, id)
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}
	if pet == nil {
		return nil, server.ErrBadRequest(fmt.Errorf("pet with id: %v not found", id))
	}
	// prepare
	pet.Name = pur.Name
	pet.DateOfBirth = pur.DateOfBirth

	err = db.PetUpdate(c, s.connDB, *pet)
	if err != nil {
		return nil, server.ErrInternalDB(err)
	}
	resp := pet.Response()
	return &resp, nil
}

func (s PetService) DeleteByID(c context.Context, id int) error {
	rowsAffected, err := db.PetDeleteByID(c, s.connDB, id)
	if err != nil {
		return server.ErrInternalDB(err)
	}
	if *rowsAffected == 0 {
		return server.ErrBadRequest(fmt.Errorf("bad pet id"))
	}
	return nil
}
