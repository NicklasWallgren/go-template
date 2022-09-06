package users

import (
	"github.com/NicklasWallgren/go-template/domain/users/models"

	"github.com/NicklasWallgren/go-template/adapters/driver/api/converters"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/users/response"
)

type UserOverviewAPIConverter interface {
	converters.APIResponseConverter[*models.SenderOverview, response.UserOverviewResponse]
}

type userOverviewAPIConverter struct{}

func NewUserOverviewAPIConverter() UserOverviewAPIConverter {
	return &userOverviewAPIConverter{}
}

func (u userOverviewAPIConverter) ResponseOf(user *models.SenderOverview) response.UserOverviewResponse {
	return response.UserOverviewResponse{Name: user.Name, Email: user.Email}
}
