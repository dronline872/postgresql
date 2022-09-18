package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type NameProductSearchHint struct {
	Id    int
	Price sql.NullFloat64
}

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
	cfg.MaxConns = 50
	cfg.MinConns = 50
	dbpool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()
	duration := time.Duration(10 * time.Second)
	threads := 1000
	fmt.Println("start attack")
	res := attack(ctx, duration, threads, dbpool)
	fmt.Println("duration:", res.Duration)
	fmt.Println("threads:", res.Threads)
	fmt.Println("queries:", res.QueriesPerformed)
	qps := res.QueriesPerformed / uint64(res.Duration.Seconds())
	fmt.Println("QPS:", qps)

}

func attack(ctx context.Context, duration time.Duration, threads int, dbpool *pgxpool.Pool) AttackResults {
	var queries uint64
	attacker := func(stopAt time.Time) {
		for {
			_, err := search(ctx, dbpool, "Товар", 5)
			if err != nil {
				log.Fatal(err)
			}
			atomic.AddUint64(&queries, 1)
			if time.Now().After(stopAt) {
				return
			}
		}
	}
	var wg sync.WaitGroup
	wg.Add(threads)
	startAt := time.Now()
	stopAt := startAt.Add(duration)
	for i := 0; i < threads; i++ {
		go func() {
			attacker(stopAt)
			wg.Done()
		}()
	}
	wg.Wait()
	return AttackResults{
		Duration:         time.Now().Sub(startAt),
		Threads:          threads,
		QueriesPerformed: queries,
	}
}

func search(ctx context.Context, dbpool *pgxpool.Pool, prefix string, limit int) ([]NameProductSearchHint, error) {
	const sql = `
	SELECT "price", "id"
	FROM "products"
	WHERE "name" LIKE $1
	LIMIT $2;
`
	pattern := prefix + "%"
	rows, err := dbpool.Query(ctx, sql, pattern, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query data: %w", err)
	}
	// Вызов Close нужен, чтобы вернуть соединение в пул
	defer rows.Close()
	// В слайс hints будут собраны все строки, полученные из базы
	var hints []NameProductSearchHint
	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		var hint NameProductSearchHint
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
