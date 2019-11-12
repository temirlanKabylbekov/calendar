package interfaces

import "github.com/temirlanKabylbekov/calendar/internal/models"

type UserRepository interface {
	Store(u *models.User) (models.UserId, error)
	Find(id models.UserId) (*models.User, error)
	Delete(id models.UserId) error
}

type CalendarRepository interface {
}

type EventRepository interface {
}

type EventOccurrenceRepository interface {
}
