package users

import (
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/converters"
	apiError "github.com/NicklasWallgren/go-template/adapters/driven/api/errors"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/request"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	userResponse "github.com/NicklasWallgren/go-template/adapters/driven/api/users/response"
	"github.com/NicklasWallgren/go-template/adapters/driver/persistence/models"
	services "github.com/NicklasWallgren/go-template/domain/users"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	"github.com/NicklasWallgren/go-template/infrastructure/logger"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service      services.UserService
	logger       logger.Logger
	apiConverter UserApiConverter
	validator    binding.StructValidator
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger logger.Logger, apiConverter UserApiConverter, validator binding.StructValidator) UserController {
	return UserController{
		service:      userService,
		logger:       logger,
		apiConverter: apiConverter,
		validator:    validator,
	}
}

// GetOneUserById gets one user
func (u UserController) GetOneUserById(ctx *gin.Context) (response.ApiResponseEnvelop, error) {
	id, err := request.GetParamInt(ctx, "id")
	if err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	user, err := u.service.FindOneUserById(ctx.Request.Context(), uint(id))
	if err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	return response.NewApiResponseWithPayload(http.StatusOK, u.apiConverter.ResponseOf(*user)), nil
}

// FindAllUsers godoc
// @Summary 	Retrieve users paginated
// @BasePath 	/api/users
// @Success		200 {object} response.PageableResponse[userResponse.UserResponse]
// @Success		400 {object} response.ApiResponseEnvelop[userResponse.UserResponse]
// @Router 		/api/users [get]
// TODO, support for generics https://github.com/swaggo/swag/issues/1170
func (u UserController) FindAllUsers(ctx *gin.Context) (response.ApiResponseEnvelop, error) {
	pagination, err := request.Into(ctx, models.NewPaginationWithDefaults())
	if err != nil {
		return nil, err
	}

	// TODO, support for predicate/criteria?
	userPage, err := u.service.FindAllUser(ctx.Request.Context(), &pagination)
	if err != nil {
		return nil, err
	}

	// TODO, inject converter
	converter := converters.PageableResponseConverter[entities.User, userResponse.UserResponse]{}
	return response.NewApiResponseWithPayload(http.StatusOK, converter.ResponseOf(userPage, u.apiConverter)), nil
}

// SaveUser saves the user
func (u UserController) SaveUser(ctx *gin.Context) (response.ApiResponseEnvelop, error) {
	request, err := request.IntoAndValidate(ctx, u.validator, CreateUserRequest{})
	if err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithError(err))
	}

	persistedUser, err := u.service.CreateUser(ctx.Request.Context(), u.apiConverter.EntityOf(request))
	if err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithError(err))
	}

	return response.NewApiResponseWithPayload(http.StatusCreated, u.apiConverter.ResponseOf(*persistedUser)), nil
}

// UpdateUser updates user
func (u UserController) UpdateUser(ctx *gin.Context) (response.ApiResponseEnvelop, error) {
	request, err := request.IntoAndValidate(ctx, u.validator, UpdateUserRequest{})
	if err != nil {
		return nil, err
	}

	toBeUpdatedUser, err := u.apiConverter.UpdatedEntityOf(ctx.Request.Context(), request)
	if err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithStatusAndMessageAndError(http.StatusInternalServerError, "unable to update user, please try again", err))
	}

	persistedUser, err := u.service.UpdateUser(ctx.Request.Context(), toBeUpdatedUser)
	if err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithError(err))
	}

	return response.NewApiResponseWithPayload(http.StatusOK, u.apiConverter.ResponseOf(*persistedUser)), nil
}

// DeleteUserById deletes user
func (u UserController) DeleteUserById(ctx *gin.Context) (response.ApiResponseEnvelop, error) {
	id, err := request.GetParamInt(ctx, "id")
	if err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	if err := u.service.DeleteUserById(ctx.Request.Context(), uint(id)); err != nil {
		return nil, apiError.NewApiErrorWith(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	return response.NewApiResponseEnvelop(http.StatusNoContent), nil
}
