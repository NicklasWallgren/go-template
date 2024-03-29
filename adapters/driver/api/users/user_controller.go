package users

import (
	"net/http"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/users"

	"github.com/NicklasWallgren/go-template/adapters/driven/logger"

	"github.com/NicklasWallgren/go-template/adapters/driven/persistence/models"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/converters"
	apiError "github.com/NicklasWallgren/go-template/adapters/driver/api/errors"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/request"
	"github.com/NicklasWallgren/go-template/adapters/driver/api/response"
	userResponse "github.com/NicklasWallgren/go-template/adapters/driver/api/users/response"

	services "github.com/NicklasWallgren/go-template/domain/users"
	"github.com/NicklasWallgren/go-template/domain/users/entities"
	userModels "github.com/NicklasWallgren/go-template/domain/users/models"
	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

// UserController is a struct which handles the typical http requests for a user.
type UserController struct {
	service              services.UserService
	logger               logger.Logger
	apiConverter         UserAPIConverter
	overviewAPIConverter UserOverviewAPIConverter
	validator            binding.StructValidator
}

// NewUserController creates a new [NewUserController].
func NewUserController(
	userService services.UserService,
	logger logger.Logger,
	apiConverter UserAPIConverter,
	overviewAPIConverter UserOverviewAPIConverter,
	validator binding.StructValidator,
) UserController {
	return UserController{
		service:              userService,
		logger:               logger,
		apiConverter:         apiConverter,
		overviewAPIConverter: overviewAPIConverter,
		validator:            validator,
	}
}

// FindOneUserByID retrieves a user by the provided ID.
// @Summary 	Retrieves a user by the provided ID.
// @Success		200 {object} response.PageableResponse[response.UserResponse]
// @Failure		400 {object} response.APIErrorResponse "in case of a bad request"
// @Failure		404 {object} response.APIErrorResponse "if an unknown ID is provided"
// @Router 		/users/{id} [get].
func (u UserController) FindOneUserByID(ctx *gin.Context) (*response.APIResponseEnvelope, error) {
	id, err := request.GetParamInt(ctx, "id")
	if err != nil {
		return nil, apiError.NewAPIError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	user, err := u.service.FindOneUserByID(ctx.Request.Context(), uint(id))
	if err != nil {
		return nil, apiError.NewAPIError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	return response.NewWithResponse(http.StatusOK, u.apiConverter.ResponseOf(user)), nil
}

// FindAllUsers retrieves paginated response of users.
// @Summary 	Retrieves paginated response of users
// @Success		200 {object} response.PageableResponse[response.UserResponse]
// @Failure		400 {object} response.APIErrorResponse "in case of an error"
// @Router 		/users [get].
func (u UserController) FindAllUsers(ctx *gin.Context) (*response.APIResponseEnvelope, error) {
	criteriaAndPagination, err := request.Into(
		ctx, models.NewCriteriaAndPagination[users.FindAllCriteria](users.FindAllCriteria{}))
	if err != nil {
		return nil, err
	}

	userPage, err := u.service.FindAllUsersByCriteria(ctx.Request.Context(), &criteriaAndPagination)
	if err != nil {
		return nil, err
	}

	// TODO, inject converter
	converter := converters.PageableResponseConverter[*entities.User, userResponse.UserResponse]{}

	return response.NewWithResponse(http.StatusOK, converter.ResponseOf(userPage, u.apiConverter)), nil
}

// Overview retrieves overview information about users.
// @Summary 	Retrieves paginated response of users
// @Success		200 {object} response.PageableResponse[response.UserResponse]
// @Failure		400 {object} response.APIErrorResponse "in case of an error"
// @Router 		/users/overview [get].
func (u UserController) Overview(ctx *gin.Context) (*response.APIResponseEnvelope, error) {
	criteriaAndPagination, err := request.Into(
		ctx, models.NewCriteriaAndPagination[users.OverviewCriteria](users.OverviewCriteria{}))
	if err != nil {
		return nil, err
	}

	userPage, err := u.service.Overview(ctx.Request.Context(), &criteriaAndPagination)
	if err != nil {
		return nil, err
	}

	// TODO, inject converter
	converter := converters.PageableResponseConverter[*userModels.SenderOverview, userResponse.UserOverviewResponse]{}

	return response.NewWithResponse(http.StatusOK, converter.ResponseOf(userPage, u.overviewAPIConverter)), nil
}

// CreateUser creates a user using the prerequisites provided.
// @Param 		request body CreateUserRequest true "query params"
// @Summary 	Creates a user using the prerequisites provided
// @Success		201 {object} response.UserResponse "if a new user was created"
// @Failure		400 {object} response.APIErrorResponse "in case of a bad request"
// @Router 		/users [post].
func (u UserController) CreateUser(ctx *gin.Context) (*response.APIResponseEnvelope, error) {
	request, err := request.IntoAndValidate(ctx, u.validator, CreateUserRequest{})
	if err != nil {
		return nil, err
	}

	persistedUser, err := u.service.CreateUser(ctx.Request.Context(), u.apiConverter.EntityOf(request))
	if err != nil {
		return nil, apiError.NewAPIError(apiError.WithError(err))
	}

	return response.NewWithResponse(http.StatusCreated, u.apiConverter.ResponseOf(persistedUser)), nil
}

// UpdateUser updates an existing user.
// @Summary 	Updates an existing user.
// @Success		200 {object} response.UserResponse "the updated users"
// @Failure		400 {object} response.APIErrorResponse "in case of a bad request"
// @Failure		500 {object} response.APIErrorResponse "in case of an internal error"
// @Router 		/users/{id} [post].
func (u UserController) UpdateUser(ctx *gin.Context) (*response.APIResponseEnvelope, error) {
	request, err := request.IntoAndValidate(ctx, u.validator, UpdateUserRequest{})
	if err != nil {
		return nil, err
	}

	toBeUpdatedUser, err := u.apiConverter.UpdatedEntityOf(ctx.Request.Context(), request)
	if err != nil {
		return nil, apiError.NewAPIError(
			apiError.WithStatusAndMessageAndError(
				http.StatusInternalServerError, "unable to update user, please try again", err))
	}

	persistedUser, err := u.service.UpdateUser(ctx.Request.Context(), toBeUpdatedUser)
	if err != nil {
		return nil, apiError.NewAPIError(apiError.WithError(err))
	}

	return response.NewWithResponse(http.StatusOK, u.apiConverter.ResponseOf(persistedUser)), nil
}

// DeleteUserByID deletes a user by id.
// @Summary 	Deletes a user by id.
// @Success		204 "if the user is deleted successfully"
// @Failure		400 {object} response.APIErrorResponse "in case of a bad request"
// @Failure		500 {object} response.APIErrorResponse "in case of an internal error"
// @Router 		/users/{id} [delete].
func (u UserController) DeleteUserByID(ctx *gin.Context) (*response.APIResponseEnvelope, error) {
	id, err := request.GetParamInt(ctx, "id")
	if err != nil {
		return nil, apiError.NewAPIError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	if err := u.service.DeleteUserByID(ctx.Request.Context(), uint(id)); err != nil {
		return nil, apiError.NewAPIError(apiError.WithStatusAndError(http.StatusBadRequest, err))
	}

	return response.NewEnvelope(http.StatusNoContent), nil
}
