package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	EventsDay   = "day"
	EventsWeek  = "week"
	EventsMount = "mount"
)

var (
	ErrCreateEvent    = errors.New("create event error")
	ErrUpdateEvent    = errors.New("update event error")
	ErrDeleteEvent    = errors.New("delete event error")
	ErrSearchEvent    = errors.New("search event error")
	ErrInvalidRequest = errors.New("invalid request")
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	UserID      int64     `json:"user_id"`
	EventDate   time.Time `json:"event_date"`
	Description string    `json:"description"`
}

func (e Event) String() string {
	v, _ := json.Marshal(e)
	return string(v)
}

type EventList []Event

func (el EventList) String() string {
	v, _ := json.Marshal(el)
	return string(v)
}

type EventStorer interface {
	CreateEvent(ctx context.Context, event Event) error
	ReadEvent(ctx context.Context, eventID uuid.UUID) (*Event, error)
	UpdateEvent(ctx context.Context, event Event) (*Event, error)
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
	SearchEvent(ctx context.Context, userID int64, startDate, endData time.Time) (EventList, error)
}

type EventStore struct {
	EventStorer
}

func NewStore(es EventStorer) *EventStore {
	return &EventStore{
		EventStorer: es,
	}
}

func (st *EventStore) CreateEvent(ctx context.Context, event Event) (*Event, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	event.ID = uuid.New()
	err := st.EventStorer.CreateEvent(ctx, event)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrCreateEvent, err)
	}

	return &event, nil
}

func (st *EventStore) UpdateEvent(ctx context.Context, event Event) (*Event, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	getEvent, err := st.EventStorer.ReadEvent(ctx, event.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrUpdateEvent.Error(), err)
	}
	if getEvent.UserID != event.UserID && event.UserID != 0 {
		getEvent.UserID = event.UserID
	}
	if getEvent.EventDate != event.EventDate &&
		event.EventDate.Unix() > time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC).Unix() {
		getEvent.EventDate = event.EventDate
	}

	if event.Description != "" {
		getEvent.Description = event.Description
	}

	updEvent, err := st.EventStorer.UpdateEvent(ctx, *getEvent)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrUpdateEvent.Error(), err)
	}

	return updEvent, nil
}

func (st *EventStore) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if _, err := st.EventStorer.ReadEvent(ctx, eventID); err != nil {
		return fmt.Errorf("%s: %w", ErrDeleteEvent.Error(), err)
	}

	if err := st.EventStorer.DeleteEvent(ctx, eventID); err != nil {
		return fmt.Errorf("%s: %w", ErrDeleteEvent.Error(), err)
	}

	return nil
}

func (st *EventStore) SearchEvent(ctx context.Context, userID int64, data time.Time, eventsFor string) (EventList, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	switch eventsFor {
	case EventsDay:
		eventsArr, err := st.EventStorer.SearchEvent(ctx, userID, data, data)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ErrSearchEvent.Error(), err)
		}
		return eventsArr, nil
	case EventsWeek:
		eventsArr, err := st.EventStorer.SearchEvent(ctx, userID, data, data.AddDate(0, 0, 7))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ErrSearchEvent.Error(), err)
		}
		return eventsArr, nil
	case EventsMount:
		eventsArr, err := st.EventStorer.SearchEvent(ctx, userID, data, data.AddDate(0, 1, 0))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ErrSearchEvent.Error(), err)
		}
		return eventsArr, nil
	}

	return nil, ErrInvalidRequest
}
