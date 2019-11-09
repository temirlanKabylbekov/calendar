package models

type EventId int

type Event struct {
	Timestamped
	Authored
	StoreDeleted

	Id          EventId
	Calendar    Calendar
	Name        string
	Description string
	Assignees   []User
}
