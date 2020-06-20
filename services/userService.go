package services

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Email    string
	Password string
	Role
}

type Role int

const (
	ADMIN Role = iota
	USER
)

type IUserRepository interface {
	Add(user User) error
	Remove(userId uuid.UUID) error
	Update(user User) (*User, error)
	Find(userId uuid.UUID) (*User, error)
	FindAll() ([]User, error)
}

type UserRepository struct {
}

type IUserService interface {
	Add(user User) error
	Remove(userId uuid.UUID) error
	Update(user User) (*User, error)
	Find(userId uuid.UUID) (*User, error)
	FindAll() ([]User, error)
}

type UserService struct {
	UserRepository
}
