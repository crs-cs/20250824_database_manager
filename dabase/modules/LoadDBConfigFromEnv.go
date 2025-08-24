package mod

import (
	"fmt"

	"github.com/joho/godotenv"
	db "shioji.cloud/app/dabase"
)

func LoadDBConfigFromEnv() (db.DBConfig, error) {
	_ = godotenv.Load() // .envがあれば読む（無くてもOK）

	cfg := db.DBConfig{
		User: Getenv("DB_USER", ""),
		Pass: Getenv("DB_PASS", ""),
		Host: Getenv("DB_HOST", "127.0.0.1"),
		Port: Getenv("DB_PORT", "3306"),
		Name: Getenv("DB_NAME", ""),
		Loc:  Getenv("DB_LOC", "Asia/Tokyo"),
	}
	if cfg.User == "" || cfg.Pass == "" || cfg.Name == "" {
		return cfg, fmt.Errorf("環境変数 DB_USER/DB_PASS/DB_NAME は必須です")
	}
	return cfg, nil
}
