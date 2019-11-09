package models

import (
	"fmt"
)

type UserNotFoundError struct {
	UserId UserId
}

func (e UserNotFoundError) Error() string {
	return fmt.Sprintf("user with given id: %v not found", e.UserId)
}
