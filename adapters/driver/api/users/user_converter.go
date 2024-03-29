package users

import (
	"context"

	"github.com/NicklasWallgren/go-template/domain/users/models"

	"github.com/NicklasWallgren/go-template/adapters/driver/api/converters"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/users/response"

	domain "github.com/NicklasWallgren/go-template/domain/users"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
)

type UserAPIConverter interface {
	converters.APIResponseConverter[*entities.User, response.UserResponse]
	converters.APIRequestCreateConverter[CreateUserRequest, entities.User]
	converters.APIRequestUpdateConverter[UpdateUserRequest, entities.User]
}

type userAPIConverter struct {
	userService domain.UserService
}

func NewUserAPIConverter(userService domain.UserService) UserAPIConverter {
	return &userAPIConverter{userService}
}

func (u userAPIConverter) ResponseOf(user *entities.User) response.UserResponse {
	return response.UserResponse{Name: user.Name, Email: user.Email}
}

func (u userAPIConverter) OverviewResponseOf(user *models.SenderOverview) response.UserOverviewResponse {
	return response.UserOverviewResponse{Name: user.Name, Email: user.Email}
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
