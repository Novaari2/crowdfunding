package user

import "time"

type RegisterUserInput struct {
	Name             string `json:"name" binding:"required"`
	Email            string `json:"email" binding:"required,email"`
	Occupation       string
	Password         string `json:"password" binding:"required"`
	Avatar_File_Name string
	Role             string
	Token            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type LoginInput struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}