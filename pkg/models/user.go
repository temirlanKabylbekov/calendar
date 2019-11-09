package models

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
