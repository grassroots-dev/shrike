package store

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/gofrs/uuid"
)

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
	ID        uuid.UUID `sql:",type:uuid"`
	Title     string
	CreatorID uuid.UUID `sql:",type:uuid"`
	Creator   *User
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
	db := pg.Connect(&pg.Options{
		User:       "tern",
		Password:   "tern",
		Database:   "tern",
		MaxRetries: 10,
		Addr:       "localhost:5432",
	})
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
	db := pg.Connect(&pg.Options{
		User:       "tern",
		Password:   "tern",
		Database:   "tern",
		MaxRetries: 10,
		Addr:       "localhost:5432",
	})
	defer db.Close()
	err := deleteSchema(db)
	if err != nil {
		return err
	}
	return nil
}

// CreateMovement creates a movement.
func CreateMovement(title string) (*uuid.UUID, error) {
	db := pg.Connect(&pg.Options{
		User:       "tern",
		Password:   "tern",
		Database:   "tern",
		MaxRetries: 10,
		Addr:       "localhost:5432",
	})
	defer db.Close()
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	user, err := uuid.FromString("b581056b-b295-477a-8700-9dec232c3641")
	if err != nil {
		return nil, err
	}
	newMovement := &Movement{
		ID:        newID,
		Title:     title,
		CreatorID: user,
	}
	err = db.Insert(newMovement)
	if err != nil {
		panic(err)
	}
	return &newID, nil
}

// DestroyMovement destroys a movement.
func DestroyMovement() error {
	db := pg.Connect(&pg.Options{
		User:       "tern",
		Password:   "tern",
		Database:   "tern",
		MaxRetries: 10,
		Addr:       "localhost:5432",
	})
	defer db.Close()
	err := deleteSchema(db)
	if err != nil {
		return err
	}
	return nil
}

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
