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

// UserController is a struct which handles the typical http requests for a user.
type UserController struct {
	service      services.UserService
	logger       logger.Logger
	apiConverter UserApiConverter
	validator    binding.StructValidator
}

// NewUserController creates a new [NewUserController].
func NewUserController(
	userService services.UserService,
	logger logger.Logger,
	apiConverter UserApiConverter,
	validator binding.StructValidator,
) UserController {
	return UserController{
		service:      userService,
		logger:       logger,
		apiConverter: apiConverter,
		validator:    validator,
	}
}

// GetOneUserByID retrieves a user by the provided ID.
// @Summary 	Retrieves a user by the provided ID.
// @Success		200 {object} response.PageableResponse[response.UserResponse]
// @Failure		400 {object} response.APIError "in case of a bad request"
// @Failure		404 {object} response.APIError "if an invalid ID is provided"
// @Router 		/users/{id} [get].
func (u UserController) GetOneUserByID(ctx *gin.Context) (response.APIResponseEnvelop, error) {
	id, err := request.GetParamInt(ctx, "id")
	if err != nil {
		return nil, apiError.NewApiError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	user, err := u.service.FindOneUserByID(ctx.Request.Context(), uint(id))
	if err != nil {
		return nil, apiError.NewApiError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	return response.NewWithResponse(http.StatusOK, u.apiConverter.ResponseOf(user)), nil
}

// FindAllUsers retrieves users by paginated response
// @Summary 	Retrieve users by paginated response
// @Success		200 {object} response.PageableResponse[response.UserResponse]
// @Failure		400 {object} response.APIError "in case of an error"
// @Router 		/users [get].
func (u UserController) FindAllUsers(ctx *gin.Context) (response.APIResponseEnvelop, error) {
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
	converter := converters.PageableResponseConverter[*entities.User, userResponse.UserResponse]{}

	return response.NewWithResponse(http.StatusOK, converter.ResponseOf(userPage, u.apiConverter)), nil
}

// SaveUser creates a user using the prerequisites provided.
// @Summary 	Creates a user using the prerequisites provided
// @Success		201 {object} response.UserResponse "if a new user was created"
// @Failure		400 {object} response.APIError "in case of a bad request"
// @Router 		/users [post].
func (u UserController) SaveUser(ctx *gin.Context) (response.APIResponseEnvelop, error) {
	request, err := request.IntoAndValidate(ctx, u.validator, CreateUserRequest{})
	if err != nil {
		return nil, apiError.NewApiError(apiError.WithError(err))
	}

	persistedUser, err := u.service.CreateUser(ctx.Request.Context(), u.apiConverter.EntityOf(request))
	if err != nil {
		return nil, apiError.NewApiError(apiError.WithError(err))
	}

	return response.NewWithResponse(http.StatusCreated, u.apiConverter.ResponseOf(persistedUser)), nil
}

// UpdateUser updates an existing user.
// @Summary 	Updates an existing user.
// @Success		200 {object} response.UserResponse "the updated users"
// @Failure		400 {object} response.APIError "in case of a bad request"
// @Failure		500 {object} response.APIError "in case of an internal error"
// @Router 		/users/{id} [post].
func (u UserController) UpdateUser(ctx *gin.Context) (response.APIResponseEnvelop, error) {
	request, err := request.IntoAndValidate(ctx, u.validator, UpdateUserRequest{})
	if err != nil {
		return nil, err
	}

	toBeUpdatedUser, err := u.apiConverter.UpdatedEntityOf(ctx.Request.Context(), request)
	if err != nil {
		return nil, apiError.NewApiError(
			apiError.WithStatusAndMessageAndError(
				http.StatusInternalServerError, "unable to update user, please try again", err))
	}

	persistedUser, err := u.service.UpdateUser(ctx.Request.Context(), toBeUpdatedUser)
	if err != nil {
		return nil, apiError.NewApiError(apiError.WithError(err))
	}

	return response.NewWithResponse(http.StatusOK, u.apiConverter.ResponseOf(persistedUser)), nil
}

// DeleteUserByID deletes a user by id.
// @Summary 	Deletes a user by id.
// @Success		204 "if the user is deleted successfully"
// @Failure		400 {object} response.APIError "in case of a bad request"
// @Failure		500 {object} response.APIError "in case of an internal error"
// @Router 		/users/{id} [delete].
func (u UserController) DeleteUserByID(ctx *gin.Context) (response.APIResponseEnvelop, error) {
	id, err := request.GetParamInt(ctx, "id")
	if err != nil {
		return nil, apiError.NewApiError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	if err := u.service.DeleteUserByID(ctx.Request.Context(), uint(id)); err != nil {
		return nil, apiError.NewApiError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	return response.New(http.StatusNoContent), nil
}
