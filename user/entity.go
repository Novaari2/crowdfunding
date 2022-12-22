package user

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	Occupation     string
	Password       string
	Avatar_File_Name string
	Role           string
	Token          string
	CreatedAt      time.Time
	UpdatedAt 	   time.Time
}

