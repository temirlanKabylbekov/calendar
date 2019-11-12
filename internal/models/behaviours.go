package models

import "time"

type Timestamped struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Authored struct {
	Author User
}

type StoreDeleted struct {
	Deleted bool
}

type TimeRanged struct {
	StartAt time.Time
	EndAt   time.Time
}
