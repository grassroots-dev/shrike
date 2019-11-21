package store

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// Configuration is a struct for the Configuration model.
type Configuration struct {
	ID     uuid.UUID `sql:",type:uuid"`
	Active bool
}

func (u Configuration) String() string {
	return fmt.Sprintf("Configuration<%d %t>", u.ID, u.Active)
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
