// Package store is responsible for persisting data to postgres.
package store

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/gofrs/uuid"
)

var pgOptions = &pg.Options{
	User:       "tern",
	Password:   "tern",
	Database:   "tern",
	MaxRetries: 10,
	Addr:       "localhost:5432",
}

// Configuration is a struct for the Configuration model.
type Configuration struct {
	ID     uuid.UUID `sql:",type:uuid"`
	Active bool
}

func (u Configuration) String() string {
	return fmt.Sprintf("User<%d %t>", u.ID, u.Active)
}

// User is a struct for the user model.
type User struct {
	ID     uuid.UUID `sql:",type:uuid"`
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Emails)
}

// Movement is a core model representing an organization.
type Movement struct {
	ID            uuid.UUID `sql:",type:uuid"`
	Title         string
	Description   string
	URI           string
	FeaturedImage string
	CreatorID     uuid.UUID `sql:",type:uuid"`
	Creator       *User
}

func (s Movement) String() string {
	return fmt.Sprintf("Movement<%d %s %s>", s.ID, s.Title, s.Creator)
}

// LandingPage is a struct representing a reachable html endpoint.
type LandingPage struct {
	ID         uuid.UUID `sql:",type:uuid"`
	Title      string
	CreatorID  int64
	Creator    *User
	MovementID uuid.UUID `sql:",type:uuid"`
	Movement   *Movement
}

func (s LandingPage) String() string {
	return fmt.Sprintf("LandingPage<%d %s %s>", s.ID, s.Title, s.Creator)
}

// Story is a struct for the story model.
type Story struct {
	ID       uuid.UUID `sql:",type:uuid"`
	Title    string
	AuthorID uuid.UUID `sql:",type:uuid"`
	Author   *User
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.ID, s.Title, s.Author)
}

// InitializeDB runs the pg ORM example.
func InitializeDB() error {
	db := pg.Connect(pgOptions)
	defer db.Close()
	err := createSchema(db)
	if err != nil {
		return err
	}
	newID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user1 := &User{
		ID:     newID,
		Name:   "admin",
		Emails: []string{"stephen@callsignmedia.com"},
	}
	err = db.Insert(user1)
	if err != nil {
		return err
	}

	return nil
}

// DestroyDB runs the pg ORM example.
func DestroyDB() error {
	db := pg.Connect(pgOptions)
	defer db.Close()
	err := deleteSchema(db)
	if err != nil {
		return err
	}
	return nil
}

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

// GetMovement returns a single movement by primary key.
func GetMovement(id string) (*Movement, error) {
	db := pg.Connect(pgOptions)
	defer db.Close()

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

// a function to return columns which have been updated to insert in DB, ignoring unchanged columns.
// Uses introspection to find struct fields that should be included in update.
// TODO: This currently returns all editable fields on the type. This will result in empty values
// overwriting data.
func getColumnsToEdit() []string {
	return []string{"title", "description", "featured_image", "uri"}
}

// UpdateMovement takes in a Movement ID and updates the fields found to be non default in struct.
func UpdateMovement(updatedMovement *Movement) (*Movement, error) {
	db := pg.Connect(pgOptions)
	defer db.Close()
	_, err := db.Model(updatedMovement).Column(getColumnsToEdit()...).WherePK().Returning("*").Update(updatedMovement)
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

// A way to destry and create the existing schema easily during development.
// TODO: Code Gen this from the model structs.
func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Configuration)(nil), (*User)(nil), (*Movement)(nil), (*LandingPage)(nil), (*Story)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Configuration)(nil), (*User)(nil), (*Movement)(nil), (*LandingPage)(nil), (*Story)(nil)} {
		err := db.DropTable(model, &orm.DropTableOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
