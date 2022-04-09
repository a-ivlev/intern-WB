package cache

import (
	"context"
	"database/sql"
	"intern-WB/l0/backend/internal/app/model"
	"intern-WB/l0/backend/internal/app/repository"
	"sync"
)

type CacheInMemoryDB struct {
	*sync.Mutex
	m map[string]model.Order
}

func NewCacheDB() *CacheInMemoryDB {
	return &CacheInMemoryDB{
		Mutex: &sync.Mutex{},
		m:     make(map[string]model.Order),
	}
}

var _ repository.OrdersStorer = &CacheInMemoryDB{}

func(db *CacheInMemoryDB) CreateOrder(ctx context.Context, order model.Order) (string, error) {
		select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	db.Lock()
	defer db.Unlock()

		db.m[order.OrderUID] = order
	_, ok := db.m[order.OrderUID]
	if !ok {
		return "", sql.ErrConnDone
	}

	return order.OrderUID, nil
}

func(db *CacheInMemoryDB) GetOrder(ctx context.Context, orderUID string) (*model.Order, error) {
		select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	db.Lock()
	defer db.Mutex.Unlock()

	getOrder, ok := db.m[orderUID]
	if !ok {
		return nil, sql.ErrNoRows
	}

	return &getOrder, nil
}