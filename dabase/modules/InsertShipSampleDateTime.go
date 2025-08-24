package mod

import (
	"context"
	"database/sql"
	"log"
	"time"
)

// DATETIME/TIMESTAMP 型の列に挿入（CreatedAt が DATETIME/TIMESTAMP）
func InsertShipSampleDateTime(ctx context.Context, db *sql.DB, sample string) (lastID int64, err error) {

	const q = `INSERT INTO ship_database (CreatedAt, Sample) VALUES (?, ?)`
	createdAt := sql.NullTime{Time: time.Now(), Valid: true}

	res, err := db.ExecContext(ctx, q, createdAt, sample)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetDBTime() sql.NullTime {
	// 日本時間の Location をロード
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal("LoadLocation error:", err)
	}

	// 現在の日本時刻
	jstNow := time.Now().In(loc)

	// sql.NullTime でラップ
	createdAt := sql.NullTime{Time: jstNow, Valid: true}

	return createdAt
}
