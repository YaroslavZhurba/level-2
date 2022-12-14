package store

import (
	"dev11/src/model"
	"errors"
	"sync"
	"time"
)

type EventRepositoryImpl struct {
	store *Store
	mutex *sync.Mutex
	events map[int]model.Event
}

func (repository *EventRepositoryImpl) InsertEvent(e model.Event) error {
	repository.mutex.Lock()
	repository.events[e.ID] = e
	repository.mutex.Unlock()
	return nil
}

func (repository *EventRepositoryImpl) UpdateEvent(e model.Event) error {
	_, ok := repository.events[e.ID]
	if !ok {
		return errors.New("event not found")
	}
	repository.mutex.Lock()
	repository.events[e.ID] = e
	repository.mutex.Unlock()
	return nil
}

func (repository *EventRepositoryImpl) DeleteEvent(id int) error {
	_, ok := repository.events[id]
	if !ok {
		return errors.New("event not found")
	}
	repository.mutex.Lock()
	delete(repository.events, id)
	repository.mutex.Unlock()
	return nil
}

func (repository *EventRepositoryImpl) GetEventsForDay(userID int, date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	repository.mutex.Lock()

	for _, event := range repository.events {
		if event.Date.Day() == date.Day() && 
		  event.Date.Month() == date.Month() &&
		  event.Date.Year() == date.Year() && userID == event.UserID {
			events = append(events, event)
		}
	}
	repository.mutex.Unlock()
	return events, nil
}

func (repository *EventRepositoryImpl) GetEventsForWeek(userID int, date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	repository.mutex.Lock()

	for _, event := range repository.events {
		EventY, EventW := event.Date.ISOWeek()
		DateY, DateW := date.ISOWeek()
		if EventY == DateY && EventW == DateW && userID == event.UserID {
			events = append(events, event)
		}
	}

	repository.mutex.Unlock()
	return events, nil
}

func (repository *EventRepositoryImpl) GetEventsForMonth(userID int, date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	repository.mutex.Lock()

	for _, event := range repository.events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.UserID == userID {
			events = append(events, event)
		}
	}
	repository.mutex.Unlock()
	return events, nil
}
