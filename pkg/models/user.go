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
	fullname := fmt.Sprintf("%s %s %s", u.LastName, u.FirstName, u.MiddleName)
	fullname = strings.TrimSpace(fullname)
	fullname = strings.ReplaceAll(fullname, "  ", " ")
	return fullname
}
