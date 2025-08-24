package main

import (
	"log"
	"time"

	mod "shioji.cloud/app/dabase/modules"
)

/* ========== 利用例（main） ========== */

func main() {
	//--- 設定ロード
	conf, err := mod.LoadDBConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	//--- MySQLに接続
	db, err := mod.OpenMySQL(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("DB接続OK")

	//--- 自動キャンセル設定
	ctx, cancel := mod.Ctx(3 * time.Second)
	defer cancel()

	//--- データの挿入
	id, err := mod.InsertShipSampleDateTime(ctx, db, "SHIOJI MARU IV")
	if err != nil {
		log.Fatal("INSERT error:", err)
	}
	log.Println("INSERT OK, LastInsertID =", id)

	//--- データの取得
	rec, err := mod.GetShipRecordBySample(ctx, db, "SHIOJI MARU IV")
	if err != nil {
		log.Fatal("SELECT error:", err)
	}
	if rec == nil {
		log.Println("レコードは見つかりませんでした")
	} else {
		log.Printf("ID=%d, Sample=%s, CreatedAt=%v", rec.ID, rec.Sample, rec.CreatedAt)
	}

}
