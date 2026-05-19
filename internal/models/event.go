package models

import "time"

type Event struct {
	ID       int
	PlayerID int
	Time     time.Time
	Extra    string
}

func NewEvent(id int, playerID int, t time.Time, extra string) *Event {
	return &Event{
		ID:       id,
		PlayerID: playerID,
		Time:     t,
		Extra:    extra,
	}
}
