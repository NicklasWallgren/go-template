package event

import (
	"github.com/NicklasWallgren/go-template/adapters/driven/api/users/response"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
)

func ResponseOf(user entities.User) any {
	return response.UserResponse{Name: user.Name, Email: user.Email}
}
