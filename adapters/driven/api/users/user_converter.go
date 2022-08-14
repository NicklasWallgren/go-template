package users

import (
	"context"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/converters"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/users/response"
	domain "github.com/NicklasWallgren/go-template/domain/users"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
)

type UserApiConverter interface {
	converters.ApiResponseConverter[entities.User, response.UserResponse]
	converters.APIRequestCreateConverter[CreateUserRequest, entities.User]
	converters.APIRequestUpdateConverter[UpdateUserRequest, entities.User]
}

type userAPIConverter struct {
	userService domain.UserService
}

func NewUserAPIConverter(userService domain.UserService) UserApiConverter {
	return &userAPIConverter{userService}
}

func (u userAPIConverter) ResponseOf(user entities.User) response.UserResponse {
	return response.UserResponse{Name: user.Name, Email: user.Name}
}

func (u userAPIConverter) EntityOf(request *CreateUserRequest) entities.User {
	return entities.NewUser(request.Name, request.Email, request.Age, request.Birthday.ToTime())
}

func (u userAPIConverter) UpdatedEntityOf(ctx context.Context, request *UpdateUserRequest) (*entities.User, error) {
	user, err := u.userService.FindOneUserByID(ctx, request.ID) // nolint:wrapcheck
	if err != nil {
		return nil, err
	}

	clonedUser := *user
	clonedUser.Name = request.Name
	clonedUser.Age = request.Age
	clonedUser.Birthday = request.Birthday.ToTime()

	return &clonedUser, nil
}
