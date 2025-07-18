package main

import (
	"crud/controller"
	"crud/model"
	"flag"
	"fmt"
	"net/http"
)

var pm = model.CreatePaperModel()
var pc = controller.CreatePaperController(pm)
var ro = controller.CreateRouter(pc)

func migrate() {
	// テスト用のダミーデータを追加
	sql := `INSERT INTO papers(id, title, authors, status) VALUES
	        ('01D3XZ3ZHCP3KG9VT4FGAD8KDR', 'Deep Learning', 'Ian Goodfellow', 'Reviewed'),
	        ('01D3XZ3ZHCP3KG9VT4FGAD8KE2', 'Computer Vision', 'Jan Koenderink', 'Reviewed'),
	        ('01D3XZ3ZHCP3KG9VT4FGAD8KE3', 'Quantum Computing', 'Michael Nielsen', 'Published');`

	_, err := model.Db.Exec(sql)
	if err != nil {
		fmt.Println("Error during migration:", err)
		return
	}
	fmt.Println("Migration is successful!")
}

func main() {
	f := flag.String("option", "", "migrate database or not")
	flag.Parse()
	if *f == "migrate" {
		migrate()
	}

	// HTTPエンドポイントを論文関連の操作に更新
	http.HandleFunc("/fetch-papers", ro.FetchPaper)
	http.HandleFunc("/add-paper", ro.AddPaper)
	http.HandleFunc("/delete-paper", ro.DeletePaper)
	http.HandleFunc("/update-paper-status", ro.ExpandGraph)
	http.ListenAndServe(":8080", nil)
}
