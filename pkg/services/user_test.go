package services

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/temirlanKabylbekov/calendar/pkg/models"
	"github.com/temirlanKabylbekov/calendar/pkg/repositories"
)

func TestServiceForUserInMemoryRepository_Create(t *testing.T) {
	repository := repositories.NewUserInMemoryRepository(make(map[models.UserId]*models.User))
	service := NewUserService(repository)

	userId1, err := service.Store(&models.User{UserName: "otus", Email: "otus@otus.com"})
	require.Nil(t, err)
	require.Equal(t, models.UserId(1), userId1)

	userId2, err := service.Store(&models.User{UserName: "suto", Email: "suto@otus.com"})
	require.Nil(t, err)
	require.Equal(t, models.UserId(2), userId2)
}

func TestServiceForUserInMemoryRepository_Update(t *testing.T) {
	database := make(map[models.UserId]*models.User)
	database[models.UserId(1)] = &models.User{UserName: "otus", Email: "otus@otus.com"}

	repository := repositories.NewUserInMemoryRepository(database)
	service := NewUserService(repository)

	userId, err := service.Store(&models.User{Id: models.UserId(1), UserName: "new_otus", Email: "otus@otus.com"})
	require.Nil(t, err)
	require.Equal(t, models.UserId(1), userId)

	updatedUser, err := service.Find(models.UserId(1))
	require.Nil(t, err)
	require.Equal(t, updatedUser.UserName, "new_otus")
}

func TestServiceForUserInMemoryRepository_UpdateNotExistingUser(t *testing.T) {
	database := make(map[models.UserId]*models.User)
	database[models.UserId(1)] = &models.User{UserName: "otus", Email: "otus@otus.com"}

	repository := repositories.NewUserInMemoryRepository(database)
	service := NewUserService(repository)

	notExistingUser := &models.User{Id: models.UserId(11), UserName: "suto", Email: "suto@otus.com"}

	_, err := service.Store(notExistingUser)
	require.EqualError(t, err, "user with given id: 11 not found")
}

func TestServiceForUserInMemoryRepository_Delete(t *testing.T) {
	database := make(map[models.UserId]*models.User)
	database[models.UserId(1)] = &models.User{UserName: "otus", Email: "otus@otus.com"}

	repository := repositories.NewUserInMemoryRepository(database)
	service := NewUserService(repository)

	err := service.Delete(models.UserId(1))
	require.Nil(t, err)

	_, err = service.Find(models.UserId(1))
	require.EqualError(t, err, "user with given id: 1 not found")
}

func TestServiceForUserInMemoryRepository_DeleteNotExistingUser(t *testing.T) {
	database := make(map[models.UserId]*models.User)
	database[models.UserId(1)] = &models.User{UserName: "otus", Email: "otus@otus.com"}

	repository := repositories.NewUserInMemoryRepository(database)
	service := NewUserService(repository)

	err := service.Delete(models.UserId(11))
	require.EqualError(t, err, "user with given id: 11 not found")
}
