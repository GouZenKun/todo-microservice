package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/aarondl/sqlboiler/v4/boil"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct{}

// init()は必ず最初に実行されるメソッド
func init() {
	mysql, exists := os.LookupEnv("MYSQL_HOST")
	if !exists {
		mysql = "localhost"
	}
	dsn := "root:root@tcp(" + mysql + ":3306)/db_todo?charset=utf8mb4&parseTime=true"
	con, err := sql.Open("mysql", dsn) //DBコネクションを設定
	if err != nil {
		log.Fatal(err)
	}
	// defer con.Close()

	// データベース接続のテスト
	if err = con.Ping(); err != nil {
		log.Fatal(err)
	}

	boil.SetDB(con) //コネクション情報をグローバル領域に保存
}
