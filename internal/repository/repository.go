package repository

import (
	"context"

	"github.com/dexxp/L0/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)


type PostgresRepository struct {
	poolDb *pgxpool.Pool
}

func NewPostgresRepository(poolDb *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{poolDb: poolDb}
}

func (repo *PostgresRepository) CreateOrdersTable() error {
    _, err := repo.poolDb.Exec(context.Background(), `
        CREATE TABLE IF NOT EXISTS orders (
            order_uid VARCHAR(255) PRIMARY KEY,
            track_number VARCHAR(255),
            entry VARCHAR(255),
            delivery_info JSONB,
            payment_info JSONB,
            items JSONB,
            locale VARCHAR(255),
            date_created TIMESTAMP,
            oof_shard VARCHAR(255)
        )
    `)
    return err
}

func (repo *PostgresRepository) InsertOrder(order models.Order) error {
    _, err := repo.poolDb.Exec(context.Background(), `
        INSERT INTO orders (order_uid, track_number, entry, delivery_info, payment_info, items, locale, date_created, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    `, order.OrderUid, order.TrackNumber, order.Entry, order.Delivery, order.Payment, order.Items, order.Locale, order.DateCreated, order.OofShard)

    return err
}

func (repo *PostgresRepository) GetOrder(orderUid string) (models.Order, error) {
    var order models.Order

    err := repo.poolDb.QueryRow(context.Background(), `
        SELECT order_uid, track_number, entry, delivery_info, payment_info, items, locale, date_created, oof_shard
        FROM orders
        WHERE order_uid = $1
    `, orderUid).Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Delivery, &order.Payment, &order.Items, &order.Locale, &order.DateCreated, &order.OofShard)
    if err != nil {
        return models.Order{}, err
    }

    return order, nil
}

func (repo *PostgresRepository) GetAllOrders() ([]models.Order, error) {
    var orders []models.Order

    rows, err := repo.poolDb.Query(context.Background(), `
        SELECT order_uid, track_number, entry, delivery_info, payment_info, items, locale, date_created, oof_shard
        FROM orders
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var order models.Order
        err := rows.Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Delivery, &order.Payment, &order.Items, &order.Locale, &order.DateCreated, &order.OofShard)
        if err != nil {
            return nil, err
        }

        orders = append(orders, order)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return orders, nil
}
