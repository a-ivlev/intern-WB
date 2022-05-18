package inMemoryDB

import (
	"context"
	"database/sql"
	"l2/develop/dev11/internal/app/event"
	"sync"
	"time"

	"github.com/google/uuid"
)

var _ event.EventStorer = &InMemoryDB{}

type InMemoryDB struct {
	sync.Mutex
	eventTable map[uuid.UUID]event.Event
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{

		eventTable: make(map[uuid.UUID]event.Event, 100),
	}
}

func (db *InMemoryDB) CreateEvent(ctx context.Context, event event.Event) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	db.eventTable[event.ID] = event

	return nil
}

func (db *InMemoryDB) ReadEvent(ctx context.Context, eventID uuid.UUID) (*event.Event, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	getEvent, ok := db.eventTable[eventID]
	if !ok {
		return nil, sql.ErrNoRows
	}

	return &getEvent, nil
}

func (db *InMemoryDB) UpdateEvent(ctx context.Context, event event.Event) (*event.Event, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	_, ok := db.eventTable[event.ID]
	if !ok {
		return nil, sql.ErrNoRows
	}

	db.eventTable[event.ID] = event

	return &event, nil
}

func (db *InMemoryDB) DeleteEvent(ctx context.Context, eventID uuid.UUID) error {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	_, ok := db.eventTable[eventID]
	if !ok {
		return sql.ErrNoRows
	}

	delete(db.eventTable, eventID)

	return nil
}

func (db *InMemoryDB) SearchEvent(ctx context.Context, userID int64, startDate, endData time.Time) (event.EventList, error) {
	db.Mutex.Lock()
	defer db.Mutex.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	eventsList := make(event.EventList, 0, 100)

	for _, elem := range db.eventTable {
		if elem.UserID == userID && startDate.Unix() <= elem.EventDate.Unix() && elem.EventDate.Unix() <= endData.Unix() {
			eventsList = append(eventsList, elem)
		}
	}

	return eventsList, nil
}
