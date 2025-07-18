package model

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error

	// データベース接続文字列を環境変数から取得
	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))
	Db, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	// データベースへの接続を確認
	err = Db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// papersテーブルが存在しない場合は作成
	sql := `CREATE TABLE IF NOT EXISTS papers(
			id varchar(26) NOT NULL,
			title varchar(255) NOT NULL,
			authors varchar(255),
			status varchar(100) NOT NULL,
			PRIMARY KEY (id)
		)`

	_, err = Db.Exec(sql)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Database connection has been established and table is set up!")
}
