package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PG struct {
	dbpool *pgxpool.Pool
}

func NewPG(dbpool *pgxpool.Pool) *PG {
	return &PG{dbpool}
}

type Product struct {
	Id    int
	Price sql.NullFloat64
}

func (s *PG) SearchProductByName(ctx context.Context, prefix string, limit int) ([]Product, error) {
	const sql = `
	SELECT "id", "price"
	FROM "products"
	WHERE "name" LIKE $1
	LIMIT $2;
`
	pattern := prefix + "%"
	rows, err := s.dbpool.Query(ctx, sql, pattern, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	// Вызов Close нужен, чтобы вернуть соединение в пул
	defer rows.Close()
	// В слайс hints будут собраны все строки, полученные из базы
	var hints []Product
	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		var hint Product
		// Scan записывает значения столбцов в свойства структуры hint
		err = rows.Scan(&hint.Id, &hint.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		hints = append(hints, hint)
	}
	// Проверка, что во время выборки данных не происходило ошибок
	if rows.Err() != nil {
		return nil, fmt.Errorf("failed to read response: %w", rows.Err())
	}
	return hints, nil
}

func (s *PG) GetProductPriceById(ctx context.Context, id int) (*sql.NullFloat64, error) {
	const sql = `
	SELECT price FROM products WHERE id =$1;
`
	result := new(Product)
	err := s.dbpool.QueryRow(ctx, sql, id).Scan(&result.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}

	return &result.Price, nil
}
