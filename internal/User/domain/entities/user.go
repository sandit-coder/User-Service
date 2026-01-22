package entities

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Email     string
	FirstName string
	LastName  string
}

func NewUser(id uuid.UUID, email string, firstName string, lastName string) *User {
	return &User{
		ID:        id,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}
