package persistence

import (
	"context"
	"os"
	"testing"

	"github.com/NicklasWallgren/go-template/adapters/driver/persistence/models"
	"github.com/NicklasWallgren/go-template/adapters/driver/persistence/users"
	"github.com/NicklasWallgren/go-template/tests/factories"
	"github.com/NicklasWallgren/go-template/tests/fakers"
	. "github.com/NicklasWallgren/go-template/tests/integration"
	"github.com/NicklasWallgren/go-template/tests/integration/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

// nolint:funlen, paralleltest
func Test(t *testing.T) {
	t.Run("GivenUser_WhenCreate_ThenIsCreated", func(t *testing.T) {
		invoke := func(uf *factories.UserFactory, repository users.UserRepository) {
			toBeCreated := fakers.User().Any()
			user, err := repository.Create(context.TODO(), &toBeCreated)

			utils.AssertNilOrFail(t, err)
			assert.Equal(t, 1, int(user.ID))
		}

		Runner(t, invoke, WithPersistenceAndApplyMigration(), TruncateDatabase)
	})

	t.Run("GivenUser_WhenFindOneForUpdate_ThenIsFoundForUpdate", func(t *testing.T) {
		invoke := func(uf *factories.UserFactory, repository users.UserRepository) {
			user := utils.ValueFromSupplierOrFail(t, uf.Any)

			user, err := repository.FindOneByIDForUpdate(context.TODO(), uint(user.ID))

			utils.AssertNilOrFail(t, err)
			assert.Equal(t, 1, int(user.ID))
		}

		Runner(t, invoke, WithPersistenceAndApplyMigration(), TruncateDatabase)
	})

	t.Run("GivenUsers_WhenFindAllWithPagination_ThenIsFoundAndPaginate", func(t *testing.T) {
		invoke := func(uf *factories.UserFactory, repository users.UserRepository) {
			utils.SuccessOrFailNow(t, func() (any, error) { return uf.Many(5) })

			page, err := repository.FindAll(context.TODO(), &models.Pagination{Page: 0, Limit: 2})

			utils.AssertNilOrFail(t, err)
			assert.Equal(t, 2, page.NumberOfElements())
			assert.Equal(t, 5, page.TotalNumberOfElements)
		}

		Runner(t, invoke, WithPersistenceAndApplyMigration(), TruncateDatabase)
	})

	t.Run("GivenUsers_WhenTotalCount_ThenIsCorrectNumber", func(t *testing.T) {
		invoke := func(uf *factories.UserFactory, repository users.UserRepository) {
			utils.SuccessOrFailNow(t, func() (any, error) { return uf.Many(5) })

			numberOfRows, err := repository.Count(context.TODO())

			utils.AssertNilOrFail(t, err)
			assert.Equal(t, 5, int(numberOfRows))
		}

		Runner(t, invoke, WithPersistenceAndApplyMigration(), TruncateDatabase)
	})

	t.Run("GivenUsers_WhenDelete_ThenIsDeleted", func(t *testing.T) {
		invoke := func(uf *factories.UserFactory, repository users.UserRepository) {
			user := utils.ValueFromSupplierOrFail(t, uf.Any)

			err := repository.DeleteByID(context.TODO(), uint(user.ID))
			utils.AssertNilOrFail(t, err)

			result, err := repository.Count(context.TODO())
			utils.AssertNilOrFail(t, err)
			assert.Equal(t, 0, int(result))
		}

		Runner(t, invoke, WithPersistenceAndApplyMigration(), TruncateDatabase)
	})

	t.Run("GivenUser_WhenFindOneByEmailWithExclusiveLock_ThenUserExists", func(t *testing.T) {
		invoke := func(uf *factories.UserFactory, repository users.UserRepository) {
			user := utils.ValueFromSupplierOrFail(t, uf.Any)

			foundUser, err := repository.FindOneByEmailWithExclusiveLock(context.TODO(), user.Email)
			utils.AssertNilOrFail(t, err)

			assert.Equal(t, user.ID, foundUser.ID)
		}

		Runner(t, invoke, WithPersistenceAndApplyMigration(), TruncateDatabase)
	})
}
