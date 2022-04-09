package repository

import (
	"context"
	"intern-WB/l0/backend/internal/app/model"
)

type OrdersStorer interface {
	CreateOrder(ctx context.Context, order model.Order) (string, error)
	GetOrder(ctx context.Context, orderUID string) (*model.Order, error)
}

type OrdersRepo struct {
	OrdersStorer
}

func NewOrdersRepo(repo OrdersStorer) *OrdersRepo {
	return &OrdersRepo{
		repo,
	}
}

func (rep *OrdersRepo) CreateOrder(ctx context.Context, order model.Order) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	entityID, err := rep.OrdersStorer.CreateOrder(ctx, order)
	if err != nil {
		return "", err
	}

	return entityID, nil
}

func (rep *OrdersRepo) GetOrder(ctx context.Context, orderUID string) (*model.Order, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	order, err := rep.OrdersStorer.GetOrder(ctx, orderUID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
