package store

import (
	"dev11/src/model"
	"time"
	"sync"
)

type EventRepository interface {
	InsertEvent(e model.Event) error
	UpdateEvent(e model.Event) error
	DeleteEvent(id int) error
	GetEventsForDay(userID int, date time.Time) ([]model.Event, error)
	GetEventsForWeek(userID int, date time.Time) ([]model.Event, error)
	GetEventsForMonth(userID int, date time.Time) ([]model.Event, error)
}

type Store struct {
	eventRepository EventRepository
}

// New ...
func New() *Store {
	return &Store{}
}

 // Event ...
 func (s *Store) Event() EventRepository {
	if s.eventRepository != nil {
		return s.eventRepository
	}

	s.eventRepository = &EventRepositoryImpl{
		store: s,
		mutex: &sync.Mutex{},
		events: make(map[int]model.Event),
	}

	return s.eventRepository
}