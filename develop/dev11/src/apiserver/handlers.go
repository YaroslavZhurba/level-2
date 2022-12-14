package apiserver

import (
	"dev11/src/model"
	"net/http"
	"fmt"
	"time"
	"strconv"

)

const dateLayout = "2006-01-02"

func (s *APIServer) createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := unmarshalEvent(r)
	if err != nil {
		throwError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}

	err = s.store.Event().InsertEvent(event)
	if err != nil {
		throwError(w, http.StatusGatewayTimeout, fmt.Errorf("insert event error: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, http.StatusOK, "Event created successfully!", []model.Event{event})
}


func (s *APIServer) updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := unmarshalEvent(r)
	if err != nil {
		throwError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}

	err = s.store.Event().UpdateEvent(event)
	if err != nil {
		throwError(w, http.StatusGatewayTimeout, fmt.Errorf("update event error: %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, http.StatusOK, "Event updated successfully!", []model.Event{event})
}

func (s *APIServer) deleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	event, err := unmarshalEvent(r)
	if err != nil {
		throwError(w, http.StatusInternalServerError, fmt.Errorf("can't unmarshal event: %s", err))
		return
	}

	err = s.store.Event().DeleteEvent(event.ID)
	if err != nil {
		throwError(w, http.StatusGatewayTimeout, fmt.Errorf("insert event error: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "Event deleted successfully!", []model.Event{event})
}

func (s *APIServer) eventsForDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	required := []string{"user_id", "date"}

	v := r.URL.Query()
	values, err := validateURLValues(v, required...)
	if err != nil {
		throwError(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse(dateLayout, values["date"])
	if err != nil {
		throwError(
			w, http.StatusBadRequest, fmt.Errorf("invalid date format: %s, waited for %s", values["date"], dateLayout),
		)
		return
	}

	userID, err := strconv.Atoi(values["user_id"])
	if err != nil || userID < 1 {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %s", values["user_id"]))
		return
	}

	events, err := s.store.Event().GetEventsForDay(userID, date)
	if err != nil {
		throwError(w, http.StatusGatewayTimeout, fmt.Errorf("get events for day error: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "Got events!", events)
}

func (s *APIServer) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	required := []string{"user_id", "date"}

	v := r.URL.Query()
	values, err := validateURLValues(v, required...)
	if err != nil {
		throwError(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse(dateLayout, values["date"])
	if err != nil {
		throwError(
			w, http.StatusBadRequest, fmt.Errorf("invalid date format: %s, waited for %s", values["date"], dateLayout),
		)
		return
	}

	userID, err := strconv.Atoi(values["user_id"])
	if err != nil || userID < 1 {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %s", values["user_id"]))
		return
	}

	events, err := s.store.Event().GetEventsForWeek(userID, date)
	if err != nil {
		throwError(w, http.StatusGatewayTimeout, fmt.Errorf("get events for day error: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "Got events!", events)
}

func (s *APIServer) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid method: %v", r.Method))
		return
	}

	required := []string{"user_id", "date"}

	v := r.URL.Query()
	values, err := validateURLValues(v, required...)
	if err != nil {
		throwError(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse(dateLayout, values["date"])
	if err != nil {
		throwError(
			w, http.StatusBadRequest, fmt.Errorf("invalid date format: %s, waited for %s", values["date"], dateLayout),
		)
		return
	}

	userID, err := strconv.Atoi(values["user_id"])
	if err != nil || userID < 1 {
		throwError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID: %s", values["user_id"]))
		return
	}

	events, err := s.store.Event().GetEventsForMonth(userID, date)
	if err != nil {
		throwError(w, http.StatusGatewayTimeout, fmt.Errorf("get events for day error: %s", err))
		return
	}

	writeResponse(w, http.StatusOK, "Got events!", events)
}
