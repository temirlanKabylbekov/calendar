package models

type EventOccurrenceId int

type EventOccurrence struct {
	TimeRanged

	Id    EventOccurrenceId
	Event Event
}
