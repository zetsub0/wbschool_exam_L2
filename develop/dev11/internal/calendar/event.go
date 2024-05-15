package calendar

import "time"

type Event struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

func NewEvent(Time time.Time, Name string) *Event { // Create event struct
	LastIdMutex.Lock()
	LastID++
	LastIdMutex.Unlock()
	return &Event{LastID, Name, Time}
}
