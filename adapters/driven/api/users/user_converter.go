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
	converters.ApiRequestCreateConverter[CreateUserRequest, entities.User]
	converters.ApiRequestUpdateConverter[UpdateUserRequest, entities.User]
}

type userApiConverter struct {
	userService domain.UserService
}

func NewUserApiConverter(userService domain.UserService) UserApiConverter {
	return &userApiConverter{userService}
}

// To ensure that UserApiConverter implements the ApiResponseConverter interface
var _ converters.ApiResponseConverter[entities.User, response.UserResponse] = (*userApiConverter)(nil)

// To ensure that UserApiConverter implements the ApiRequestCreateConverter interface
var _ converters.ApiRequestCreateConverter[CreateUserRequest, entities.User] = (*userApiConverter)(nil)

// To ensure that UserApiConverter implements the ApiRequestUpdateConverter interface
var _ converters.ApiRequestUpdateConverter[UpdateUserRequest, entities.User] = (*userApiConverter)(nil)

func (u userApiConverter) ResponseOf(user entities.User) response.UserResponse {
	return response.UserResponse{Name: user.Name, Email: user.Name}
}

func (u userApiConverter) EntityOf(request *CreateUserRequest) entities.User {
	return entities.NewUser(request.Name, request.Email, request.Age, request.Birthday.ToTime())
}

func (u userApiConverter) UpdatedEntityOf(ctx context.Context, request *UpdateUserRequest) (*entities.User, error) {
	user, err := u.userService.FindOneUserById(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	clonedUser := *user
	clonedUser.Name = request.Name
	clonedUser.Age = request.Age
	clonedUser.Birthday = request.Birthday.ToTime()

	return &clonedUser, nil
}
