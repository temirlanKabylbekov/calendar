package repositories

import "github.com/temirlanKabylbekov/calendar/internal/models"

type UserInMemoryRepository struct {
	m      map[models.UserId]*models.User
	idIncr models.UserId
}

func NewUserInMemoryRepository(database map[models.UserId]*models.User) *UserInMemoryRepository {
	return &UserInMemoryRepository{m: database}
}

func (r *UserInMemoryRepository) Store(u *models.User) (models.UserId, error) {
	if u.Id == 0 {
		newId := r.idIncr.Incr()
		r.m[newId] = u
		r.idIncr = newId
		return newId, nil
	}
	if _, ok := r.m[u.Id]; !ok {
		return 0, models.UserNotFoundError{UserId: u.Id}
	}
	r.m[u.Id] = u
	return u.Id, nil
}

func (r UserInMemoryRepository) Find(id models.UserId) (*models.User, error) {
	user, ok := r.m[id]
	if !ok {
		return user, models.UserNotFoundError{UserId: id}
	}
	return user, nil
}

func (r UserInMemoryRepository) Delete(id models.UserId) error {
	if _, ok := r.m[id]; !ok {
		return models.UserNotFoundError{UserId: id}
	}
	delete(r.m, id)
	return nil
}
