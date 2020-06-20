package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Email    string
	Password string
	Role
	Score
}

type Role int

const (
	ADMIN Role = iota
	USER
)
