package models

import (
	"fmt"
	"strings"
)

type UserId int

func (uid UserId) Incr() UserId {
	return UserId(int(uid) + 1)
}

type User struct {
	Timestamped
	StoreDeleted

	Id         UserId
	UserName   string
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
}

func (u User) GetFullName() string {
	fullName := fmt.Sprintf("%s %s %s", u.LastName, u.FirstName, u.MiddleName)
	fullName = strings.TrimSpace(fullName)
	fullName = strings.ReplaceAll(fullName, "  ", " ")
	return fullName
}
