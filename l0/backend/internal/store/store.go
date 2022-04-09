package store

import (
	"context"
	"errors"
	"fmt"
	"intern-WB/l0/backend/internal/app/model"
	"intern-WB/l0/backend/internal/store/cache"
	"intern-WB/l0/backend/internal/store/pgDB"
)

type Store struct {
	*cache.CacheInMemoryDB
	*pgDB.PostgreDB
}

func NewStore(ctx context.Context, dsn string) (*Store, error) {
	pg, err := pgDB.NewPostgreDB(dsn)
	if err != nil {
		return nil, err
	}

	cache := cache.NewCacheDB()

	// Разогрев кеша.
	allOrdersPG, err := pg.GetAllOrders(ctx)
	for _, order := range allOrdersPG {
		_, err := cache.CreateOrder(ctx, order)
		if err != nil {
			return nil, fmt.Errorf("%s: %s", ErrWriteCache, err.Error())
		}
	}

	return &Store{
		cache,
		pg,
	}, nil
}

func (s *Store) CreateOrder(ctx context.Context, order model.Order) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	pgOrderID, err := s.PostgreDB.CreateOrder(ctx, order)
	if err != nil {
		return "", err
	}

	cacheOrderID, err := s.CacheInMemoryDB.CreateOrder(ctx, order)
	if err != nil {
		return "", err
	}

	if pgOrderID != cacheOrderID {
		return "", errors.New("database write error")
	}

	return cacheOrderID, nil
}

func (s *Store) GetOrder(ctx context.Context, orderUID string) (*model.Order, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	cacheOrder, err := s.CacheInMemoryDB.GetOrder(ctx, orderUID)
	if err != nil {
		return nil, err
	}

	return cacheOrder, nil
}
