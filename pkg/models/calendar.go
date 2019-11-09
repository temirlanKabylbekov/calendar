package models

type CalendarId int

type Calendar struct {
	Authored
	Timestamped
	StoreDeleted

	Id                   CalendarId
	Name                 string
	SharedWithToOnlyRead []User
	SharedWithToEditToo  []User
}
