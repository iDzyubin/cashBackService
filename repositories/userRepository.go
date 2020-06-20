package repositories

import (
	. "cashBackService/models"
	"github.com/google/uuid"
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

func (repository *UserRepository) Add(user User) error {
	panic("implement me")
}

func (repository *UserRepository) Remove(userId uuid.UUID) error {
	panic("implement me")
}

func (repository *UserRepository) Update(user User) (*User, error) {
	panic("implement me")
}

func (repository *UserRepository) Find(userId uuid.UUID) (*User, error) {
	panic("implement me")
}

func (repository *UserRepository) FindAll() ([]User, error) {
	panic("implement me")
}
