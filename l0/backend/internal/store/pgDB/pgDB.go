package pgDB

import (
	"context"
	"database/sql"
	"fmt"
	"intern-WB/l0/backend/internal/app/model"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgreDB struct {
	db *sql.DB
}

func NewPostgreDB(dsn string) (*PostgreDB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	pg := &PostgreDB{
		db: db,
	}
	return pg, nil
}

func (pg *PostgreDB) Close() {
	pg.db.Close()
}

func (pg *PostgreDB) CreateOrder(ctx context.Context, order model.Order) (string, error) {
	_, err := pg.db.ExecContext(ctx, "INSERT INTO entity (id, orders) VALUES($1, $2)", order.OrderUID, order)
	if err != nil {
		return "", err
	}

	return order.OrderUID, nil
}

func (pg *PostgreDB) GetOrder(ctx context.Context, orderUID string) (*model.Order, error) {
	entity := new(model.Entity)
	err := pg.db.QueryRowContext(ctx, "SELECT id, orders FROM entity WHERE id = $1;", orderUID).Scan(&entity.ID, &entity.Order)
	if err != nil {
		return nil, err
	}

	return &entity.Order, nil
}

func (pg *PostgreDB) GetAllOrders(ctx context.Context) ([]model.Order, error) {
	hints := make([]model.Order, 0, 100)
	rows, err := pg.db.QueryContext(ctx, "SELECT id, orders FROM entity;")
	if err != nil {
		return nil, err
	}
	// Вызов Close нужен, чтобы вернуть соединение в пул.
	defer rows.Close()

	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		entity := new(model.Entity)
		if err := rows.Scan(&entity.ID, &entity.Order); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		hints = append(hints, entity.Order)
	}

	return hints, nil
}
