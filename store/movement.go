package store

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/gofrs/uuid"
)

// CreateMovement creates a movement.
func CreateMovement(newMovement Movement) (*Movement, error) {
	db := pg.Connect(pgOptions)
	defer db.Close()

	// CREATE THE UNIQUE KEY HERE FOR MORE CONTROL
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	newMovement.ID = newID
	err = db.Insert(&newMovement)
	if err != nil {
		return nil, err
	}
	return &newMovement, nil
}

// ReadMovement returns a single movement by primary key.
func ReadMovement(id string) (*Movement, error) {
	db := pg.Connect(pgOptions)
	defer db.Close()
	x := User{Name: "hello"}
	fmt.Println(x)
	movementID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	// Select user by primary key.
	movement := &Movement{ID: movementID}
	err = db.Select(movement)
	if err != nil {
		return nil, err
	}
	return movement, nil
}

// ListMovements returns a list of Movements.
func ListMovements() (*[]Movement, error) {
	db := pg.Connect(pgOptions)
	defer db.Close()
	var movements []Movement
	err := db.Model(&movements).Select()
	if err != nil {
		return nil, err
	}
	return &movements, nil
}

// UpdateMovement takes in a Movement ID and updates the fields found to be non default in struct.
func UpdateMovement(updatedMovement *Movement) (*Movement, error) {
	db := pg.Connect(pgOptions)
	defer db.Close()
	_, err := db.Model(updatedMovement).Column(getColumnsToEdit(*updatedMovement)...).WherePK().Returning("*").Update(updatedMovement)
	if err != nil {
		return nil, err
	}
	return updatedMovement, nil
}

// DeleteMovement removes a single movement by primary key.
func DeleteMovement(id string) (*Movement, error) {
	db := pg.Connect(pgOptions)
	defer db.Close()

	movementID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}
	// Select movement to delete by primary key.
	movement := &Movement{ID: movementID}
	err = db.Delete(movement)
	if err != nil {
		return nil, err
	}
	return movement, nil
}
