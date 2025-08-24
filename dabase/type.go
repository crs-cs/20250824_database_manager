package db

import "time"

// --- DBのログイン設定の構成
type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
	Loc  string // 例: "Asia/Tokyo" / "UTC"
}

// --- DBの構造
type ShipRecord struct {
	ID        int
	CreatedAt time.Time
	Sample    string
}
