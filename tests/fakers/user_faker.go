package fakers

import (
	"time"

	"github.com/NicklasWallgren/go-template/domain/common"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/brianvoe/gofakeit/v6"
)

type UserFaker struct {
	common.Entity
	ID        uint
	Name      string
	Email     string
	Age       uint8
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	faker     *gofakeit.Faker
}

// To ensure that UserFaker implements the EntityFaker interface.
var _ EntityFaker[entities.User] = (*UserFaker)(nil)

func NewUserFaker(faker *gofakeit.Faker) *UserFaker {
	return &UserFaker{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Birthday: faker.Date(),
		faker:    faker,
	}
}

func User() *UserFaker {
	return NewUserFaker(gofakeit.New(1))
}

func UserWithFaker(faker *gofakeit.Faker) *UserFaker {
	return NewUserFaker(faker)
}

func (u UserFaker) Any() entities.User {
	return User().Create()
}

func (u UserFaker) Create() entities.User {
	return entities.User{Name: u.Name, Email: u.Email, Birthday: u.Birthday}
}
