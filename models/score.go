package models

import "github.com/google/uuid"

type Score struct {
	Id    uuid.UUID
	Score float32
}
