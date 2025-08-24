package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// chartier で接続する例
	dsn := "chartier:ss642644@tcp(153.121.70.230:3306)/cloud_database?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}	
	defer db.Close()

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(10 * time.Minute)

	ctx, cancel := context.WithTimeout(con	text.Background(), 3*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Ping error:", err)
	}
	fmt.Println("DB接続OK")

	// 挿入例
	res, err := db.ExecContext(ctx,
		`INSERT INTO ship_database (CreatedAt, sample) VALUES (?, ?)`,
		time.Now().Format("2006-01-02"),
		"Sample",
	)

	if err != nil {
		log.Fatal("INSERT error:", err)
	}
	id, _ := res.LastInsertId()
	fmt.Println("INSERT id:", id)

	// 取得例
	type Row struct {
		ID        int
		CreatedAt time.Time
		Sample    string
	}
	var r Row
	err = db.QueryRowContext(ctx,
		`SELECT id, CreatedAt, Sample
		 FROM ship_database WHERE id=?`, id).
		Scan(&r.ID, &r.CreatedAt, &r.Sample)
	if err != nil {
		log.Fatal("SELECT error:", err)
	}
	fmt.Printf("ROW: %+v\n", r)
}
