package storage_test

import (
	"context"
	"database/sql"
	"main/pkg/domain/storage"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestIntegrationSearch(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()
	testsSearchByName := []struct {
		name    string
		store   *storage.PG
		ctx     context.Context
		prefix  string
		limit   int
		prepare func(*pgxpool.Pool)
		check   func(*testing.T, []storage.Product, error)
	}{
		{
			name:   "success",
			store:  storage.NewPG(dbpool),
			ctx:    context.Background(),
			prefix: "Товар",
			limit:  5,
			prepare: func(dbpool *pgxpool.Pool) {
				// Подготовка тестовых данных
				dbpool.Exec(context.Background(), `INSERT INTO products (
					name,
					price,
					sort,
					description,
					img
				)
				VALUES
					('Товар 1', 100.0, 100, 'Большое описание товара', './images/img1.jpg'),
					('Товар 2', 200.0, 100, 'Большое описание товара', './images/img2.jpg'),
					('Товар 3', 300.0, 100, 'Большое описание товара', './images/img3.jpg'),
					('Товар 4', 400.0, 100, 'Большое описание товара', './images/img4.jpg'),
					('Товар 5', 500.0, 100, 'Большое описание товара', './images/img5.jpg'),
					('Товар 6', 600.0, 100, 'Большое описание товара', './images/img6.jpg'),
					('Товар 7', 700.0, 100, 'Большое описание товара', './images/img7.jpg'),
					('Товар 8', 800.0, 100, 'Большое описание товара', './images/img8.jpg'),
					('Товар 9', 900.0, 100, 'Большое описание товара', './images/img9.jpg');`)
			},
			check: func(t *testing.T, hints []storage.Product, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hints)
			},
		},
	}

	for _, tt := range testsSearchByName {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(dbpool)
			hints, err := tt.store.SearchProductByName(tt.ctx, tt.prefix, tt.limit)
			tt.check(t, hints, err)
		})
	}
}

func TestGetProductPriceById(t *testing.T) {
	ctx := context.Background()
	dbpool := connect(ctx)
	defer dbpool.Close()
	testsSearchByName := []struct {
		name    string
		store   *storage.PG
		ctx     context.Context
		id      int
		prepare func(*pgxpool.Pool)
		check   func(*testing.T, *sql.NullFloat64, error)
	}{
		{
			name:  "success",
			store: storage.NewPG(dbpool),
			ctx:   context.Background(),
			id:    1,
			prepare: func(dbpool *pgxpool.Pool) {
				// Подготовка тестовых данных
				dbpool.Exec(context.Background(), `INSERT INTO products (
					name,
					price,
					sort,
					description,
					img
				)
				VALUES
					('Товар 1', 100.0, 100, 'Большое описание товара', './images/img1.jpg'),
					('Товар 2', 200.0, 100, 'Большое описание товара', './images/img2.jpg'),
				`)
			},
			check: func(t *testing.T, hints *sql.NullFloat64, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, hints)
			},
		},
	}

	for _, tt := range testsSearchByName {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare(dbpool)
			hints, err := tt.store.GetProductPriceById(tt.ctx, tt.id)
			tt.check(t, hints, err)
		})
	}
}

// Соединение с экземпляром Postgres
func connect(ctx context.Context) *pgxpool.Pool {
	url := "postgres://user:1234567@localhost:5432/lesson"
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic(err)
	}

	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		panic(err)
	}

	return dbpool
}
