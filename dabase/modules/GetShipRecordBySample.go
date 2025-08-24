package mod

import (
	"context"
	"database/sql"

	db "shioji.cloud/app/dabase"
)

func GetShipRecordBySample(ctx context.Context, mod *sql.DB, sample string) (*db.ShipRecord, error) {
	const q = `SELECT id, CreatedAt, Sample FROM ship_database WHERE Sample = ? LIMIT 1`

	row := mod.QueryRowContext(ctx, q, sample)

	var rec db.ShipRecord

	if err := row.Scan(&rec.ID, &rec.CreatedAt, &rec.Sample); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // 見つからなかった場合
		}
		return nil, err
	}
	return &rec, nil
}
