package store

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// Configuration is a struct for the Configuration model.
type Configuration struct {
	ID     int64
	Active bool
}

func (u Configuration) String() string {
	return fmt.Sprintf("User<%d %t>", u.ID, u.Active)
}

// User is a struct for the user model.
type User struct {
	ID     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Emails)
}

// Movement is a core model representing an organization.
type Movement struct {
	ID        int64
	Title     string
	CreatorID int64
	Creator   *User
}

func (s Movement) String() string {
	return fmt.Sprintf("Movement<%d %s %s>", s.ID, s.Title, s.Creator)
}

// LandingPage is a struct representing a reachable html endpoint.
type LandingPage struct {
	ID         int64
	Title      string
	CreatorID  int64
	Creator    *User
	MovementID int64
	Movement   *Movement
}

func (s LandingPage) String() string {
	return fmt.Sprintf("LandingPage<%d %s %s>", s.ID, s.Title, s.Creator)
}

// Story is a struct for the story model.
type Story struct {
	ID       int64
	Title    string
	AuthorID int64
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
	user1 := &User{
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	err = db.Insert(user1)
	if err != nil {
		return err
	}

	return nil
}

// ResetDB runs the pg ORM example.
func ResetDB() error {
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
