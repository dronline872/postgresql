package main

import (
	"context"
	"fmt"
	"log"
	"main/pkg/domain/storage"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AttackResults struct {
	Duration         time.Duration
	Threads          int
	QueriesPerformed uint64
}

func main() {
	ctx := context.Background()
	url := "postgres://user:1234567@localhost:5432/lesson"
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Fatal(err)
	}
	cfg.MaxConns = 1
	cfg.MinConns = 1
	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	// Поиск по названию товара
	products, err := storage.NewPG(dbpool).SearchProductByName(ctx, "Товар", 5)
	if err != nil {
		log.Fatal(err)
	}

	for _, product := range products {
		fmt.Printf("Product id %d price:%v\n", product.Id, product.Price.Float64)
	}

	// Получение цены товара по id
	productPrice, err := storage.NewPG(dbpool).GetProductPriceById(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	if productPrice != nil {
		fmt.Printf("Product by id price:%v\n", productPrice.Float64)
	}
}
