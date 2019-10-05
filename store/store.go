package store

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// User is a struct for the user model.
type User struct {
	ID     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.ID, u.Name, u.Emails)
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

// RunPGExample runs the pg ORM example.
func RunPGExample() error {
	db := pg.Connect(&pg.Options{
		User:       "tern",
		Password:   "tern",
		Database:   "tern",
		MaxRetries: 10,
		Addr:       "db:5432",
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

	err = db.Insert(&User{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	})
	if err != nil {
		return err
	}

	story1 := &Story{
		Title:    "Cool story",
		AuthorID: user1.ID,
	}
	err = db.Insert(story1)
	if err != nil {
		return err
	}

	// Select user by primary key.
	user := &User{ID: user1.ID}
	err = db.Select(user)
	if err != nil {
		return err
	}

	// Select all users.
	var users []User
	err = db.Model(&users).Select()
	if err != nil {
		return err
	}

	// Select story and associated author in one query.
	story := new(Story)
	err = db.Model(story).
		Relation("Author").
		Where("story.ID = ?", story1.ID).
		Select()
	if err != nil {
		return err
	}

	fmt.Println(user)
	fmt.Println(users)
	fmt.Println(story)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
	return nil
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil), (*Story)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
