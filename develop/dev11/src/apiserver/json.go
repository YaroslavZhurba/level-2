package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"dev11/src/model"
)

func unmarshalEvent(r *http.Request) (model.Event, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return model.Event{}, err
	}

	var event model.Event

	err = json.Unmarshal(b, &event)
	if err != nil {
		return model.Event{}, err
	}

	if event.ID < 1 || event.UserID < 1 {
		return model.Event{}, fmt.Errorf("invalid event ID or user ID")
	}

	return event, nil
}
