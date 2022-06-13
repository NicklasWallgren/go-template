package api

import (
	"net/http"
	"os"
	"testing"

	"github.com/NicklasWallgren/go-template/adapters/driven/api/common"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/request/types"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/response"
	"github.com/NicklasWallgren/go-template/adapters/driven/api/users"
	userResponse "github.com/NicklasWallgren/go-template/adapters/driven/api/users/response"
	"github.com/NicklasWallgren/go-template/tests/factories"
	. "github.com/NicklasWallgren/go-template/tests/integration"
	"github.com/NicklasWallgren/go-template/tests/integration/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

// nolint: funlen, paralleltest
func Test(t *testing.T) {
	t.Run("GivenUsers_WhenGetOneUser_ThenMatchSnapshot", func(t *testing.T) {
		t.Parallel()

		testFunc := func(uf *factories.UserFactory, requestHandler common.RequestHandler) {
			utils.SuccessOrFailNow(t, func() (any, error) { return uf.Many(5) }) // nolint:wrapcheck

			request := utils.NewHttpRequest(t, "GET", "/api/users/1", nil)
			userResponse := userResponse.UserResponse{}
			utils.DoHttpRequestWithResponse(t, requestHandler.Gin, request, &userResponse, utils.ExpectHttpStatus(http.StatusOK))

			snaps.MatchSnapshot(t, userResponse)
		}

		// Creates a unique database based on the test func name. Allows parallel execution
		Runner(t, testFunc, WithApplicationAndApplyMigration(WithDatabaseName(t, t.Name())), InitializeMiddlewareAndRoutes)
	})

	t.Run("GivenUsers_WhenGetUsers_ThenMatchSnapshot", func(t *testing.T) {
		testFunc := func(uf *factories.UserFactory, requestHandler common.RequestHandler) {
			utils.SuccessOrFailNow(t, func() (any, error) { return uf.Many(5) }) // nolint:wrapcheck

			request := utils.NewHttpRequest(t, "GET", "/api/users/", nil)
			pageableUserResponse := response.PageableResponse[userResponse.UserResponse]{}
			utils.DoHttpRequestWithResponse(t, requestHandler.Gin, request, &pageableUserResponse, utils.ExpectHttpStatus(http.StatusOK))

			snaps.MatchSnapshot(t, pageableUserResponse)
		}

		Runner(t, testFunc, WithApplicationAndApplyMigration(), TruncateDatabase, InitializeMiddlewareAndRoutes)
	})

	t.Run("GivenValidRequest_WhenSaveUser_ThenMatchSnapshot", func(t *testing.T) {
		// We need to create a goroutine specific faker
		faker := gofakeit.New(1)

		testFunc := func(uf *factories.UserFactory, requestHandler common.RequestHandler) {
			userRequest := users.CreateUserRequest{
				Name:     faker.Name(),
				Email:    faker.Email(),
				Age:      uint8(faker.Number(18, 150)),
				Birthday: types.Date(faker.Date()),
			}

			request := utils.NewHttpRequest(t, "POST", "/api/users/", utils.EncodeToJSON(t, &userRequest))
			userResponse := userResponse.UserResponse{}
			utils.DoHttpRequestWithResponse(t, requestHandler.Gin, request, &userResponse, utils.ExpectHttpStatus(http.StatusCreated))

			snaps.MatchSnapshot(t, userResponse)
		}

		Runner(t, testFunc, WithApplicationAndApplyMigration(), TruncateDatabase, InitializeMiddlewareAndRoutes)
	})

	t.Run("GivenValidUpdateRequest_WhenUpdateUser_ThenMatchSnapshot", func(t *testing.T) {
		// We need to create a goroutine specific faker
		faker := gofakeit.New(1)

		testFunc := func(uf *factories.UserFactory, requestHandler common.RequestHandler) {
			utils.SuccessOrFailNow(t, func() (any, error) { return uf.Any() }) // nolint:wrapcheck

			updateUserRequest := users.UpdateUserRequest{
				Name:     faker.Name(),
				Age:      uint8(faker.Number(18, 150)),
				Birthday: types.Date(faker.Date()),
			}

			request := utils.NewHttpRequest(t, "POST", "/api/users/1", utils.EncodeToJSON(t, &updateUserRequest))
			userResponse := userResponse.UserResponse{}
			utils.DoHttpRequestWithResponse(t, requestHandler.Gin, request, &userResponse, utils.ExpectHttpStatus(http.StatusOK))

			snaps.MatchSnapshot(t, userResponse)
		}

		Runner(t, testFunc, WithApplicationAndApplyMigration(), TruncateDatabase, InitializeMiddlewareAndRoutes)
	})

	t.Run("GivenValidDeleteRequest_WhenDeleteUser_ThenExpectedHttpStatus", func(t *testing.T) {
		testFunc := func(uf *factories.UserFactory, requestHandler common.RequestHandler) {
			utils.SuccessOrFailNow(t, func() (any, error) { return uf.Many(5) }) // nolint:wrapcheck

			request := utils.NewHttpRequest(t, "DELETE", "/api/users/1", nil)
			utils.DoHttpRequest(t, requestHandler.Gin, request, utils.ExpectHttpStatus(http.StatusNoContent))

			// TODO, ensure that the user has been deleted in the database?
		}

		Runner(t, testFunc, WithApplicationAndApplyMigration(), TruncateDatabase, InitializeMiddlewareAndRoutes)
	})
}
