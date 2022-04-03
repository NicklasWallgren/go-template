package factories

import (
	"context"
	pUser "github.com/NicklasWallgren/go-template/adapters/driver/persistence/users"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/NicklasWallgren/go-template/tests/fakers"
	"github.com/brianvoe/gofakeit/v6"
)

type UserFactory struct {
	userRepository pUser.UserRepository
}

// To ensure that UserFactory implements the EntityFactory interface
var _ EntityFactory[entities.User, fakers.UserFaker] = (*UserFactory)(nil)

func NewUserFactory(userRepository pUser.UserRepository) *UserFactory {
	return &UserFactory{userRepository: userRepository}
}

func (u UserFactory) Any() (*entities.User, error) {
	user, err := u.Given(fakers.User().Create())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserFactory) Many(numberOfEntities int) (users []entities.User, err error) {
	faker := gofakeit.New(1)

	for i := 0; i < numberOfEntities; i++ {
		user, err := u.Given(fakers.UserWithFaker(faker).Create())
		if err != nil {
			return users, err
		}
		users = append(users, *user)
	}
	return
}

func (u UserFactory) Given(user entities.User) (*entities.User, error) {
	return u.userRepository.Create(context.TODO(), &user)
}
