package api

import (
	"net/http"
	"time"

	"github.com/gastownhall/gascity/internal/events"
)

// eventEmitRequest is the JSON body for POST /v0/events.
type eventEmitRequest struct {
	Type    string `json:"type"`
	Actor   string `json:"actor"`
	Subject string `json:"subject,omitempty"`
	Message string `json:"message,omitempty"`
}

func parseEventFilter(r *http.Request) (events.Filter, error) {
	q := r.URL.Query()
	filter := events.Filter{
		Type:  q.Get("type"),
		Actor: q.Get("actor"),
	}
	if v := q.Get("since"); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			filter.Since = time.Now().Add(-d)
		} else {
			return events.Filter{}, err
		}
	}
	return filter, nil
}

func (s *Server) listEvents(filter events.Filter) ([]events.Event, error) {
	ep := s.state.EventProvider()
	if ep == nil {
		return []events.Event{}, nil
	}
	evts, err := ep.List(filter)
	if err != nil {
		return nil, err
	}
	if evts == nil {
		evts = []events.Event{}
	}
	return evts, nil
}
