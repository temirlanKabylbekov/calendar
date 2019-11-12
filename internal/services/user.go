package services

import (
	"github.com/temirlanKabylbekov/calendar/internal/interfaces"
	"github.com/temirlanKabylbekov/calendar/internal/models"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(r interfaces.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Store(u *models.User) (models.UserId, error) {
	return s.repo.Store(u)
}

func (s *UserService) Find(id models.UserId) (*models.User, error) {
	return s.repo.Find(id)
}

func (s *UserService) Delete(id models.UserId) error {
	return s.repo.Delete(id)
}
