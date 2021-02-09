package api

type UserType string

const (
	Artist UserType = "Artist"
	Supervisor = "Supervisor"
	Production = "Production"
)

type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	Type UserType
}

