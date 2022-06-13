package users

import (
	"fmt"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/users/response"
	"github.com/NicklasWallgren/go-template/domain/events"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
)

func ResponseOf(event *events.Event) (any, error) {
	user, ok := event.Payload.(*entities.User)
	if !ok {
		return response.UserResponse{}, fmt.Errorf("could not type cast event payload as entities.User") // nolint:goerr113
	}

	return response.UserResponse{Name: user.Name, Email: user.Email}, nil
}
