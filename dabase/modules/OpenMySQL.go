package mod

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	db "shioji.cloud/app/dabase"
)

func OpenMySQL(cfg db.DBConfig) (*sql.DB, error) {
	loc, err := time.LoadLocation(cfg.Loc)
	if err != nil {
		log.Printf("警告: time.LoadLocation(%q) 失敗: %v -> UTCにフォールバックします", cfg.Loc, err)
		loc = time.UTC
	}

	mc := mysql.NewConfig()
	mc.Net = "tcp"
	mc.Addr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	mc.User = cfg.User
	mc.Passwd = cfg.Pass
	mc.DBName = cfg.Name
	mc.Loc = loc
	mc.ParseTime = true            // DATETIME/TIMESTAMPをtime.Timeに
	mc.Params = map[string]string{ // 文字コードなど
		"charset": "utf8mb4",
	}

	db, err := sql.Open("mysql", mc.FormatDSN())
	if err != nil {
		return nil, err
	}

	// コネクションプール設定は用途に応じて調整
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(10 * time.Minute)

	// 軽い疎通確認
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("Ping error: %w", err)
	}
	return db, nil
}
