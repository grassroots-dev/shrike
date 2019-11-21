package store

import (
	"fmt"
	"reflect"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/gofrs/uuid"
	"github.com/iancoleman/strcase"
)

var pgOptions = &pg.Options{
	User:       "tern",
	Password:   "tern",
	Database:   "tern",
	MaxRetries: 10,
	Addr:       "localhost:5432",
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

// a function to return columns which have been updated to insert in DB, ignoring unchanged columns.
// Uses introspection to find struct fields that should be included in update and return as snake case table names.
func getColumnsToEdit(m Movement) []string {
	v := reflect.ValueOf(m)
	typeOfM := v.Type()
	fieldsToUpdate := []string{}
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() {
			fmt.Printf("Updating %s to %v\n", strcase.ToSnake(typeOfM.Field(i).Name), v.Field(i).Interface())

			fieldsToUpdate = append(fieldsToUpdate, strcase.ToSnake(typeOfM.Field(i).Name))
		}
	}
	return fieldsToUpdate
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
