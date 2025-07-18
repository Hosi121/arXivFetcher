package model

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
)

type PaperModel interface {
    FetchPaper(r *http.Request) (*Paper, error)
    AnalyzeCitations(r *http.Request) ([]Citation, error)
    ExpandGraph(r *http.Request) (*Graph, error)
    DeletePaper(r *http.Request) (sql.Result, error)
}

type Paper struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Status  string `json:"status"`
}

type Citation struct {
    // ここに引用情報のフィールドを定義
}

type Graph struct {
    // ここにグラフ構造のフィールドを定義
}

func CreatePaperModel() PaperModel {
	return &paperModel{}
}

func (pm *paperModel) FetchPapers() ([]*Paper, error) {
	sql := `SELECT id, title, authors, status FROM papers`

	rows, err := Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var papers []*Paper

	for rows.Next() {
		var (
			id, title, authors, status string
		)

		if err := rows.Scan(&id, &title, &authors, &status); err != nil {
			return nil, err
		}

		papers = append(papers, &Paper{
			Id:      id,
			Title:   title,
			Authors: authors,
			Status:  status,
		})
	}

	return papers, nil
}

func (pm *paperModel) AnalyzeCitations(r *http.Request) ([]Citation, error) {
    return nil, nil // Todo
}

func (pm *paperModel) ExpandGraph(r *http.Request) (*Graph, error) {
    return nil, nil // Todo
}

func (pm *paperModel) AddPaper(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	req := Paper{
		Id:      id.String(),
		Title:   r.FormValue("title"),
		Authors: r.FormValue("authors"),
		Status:  r.FormValue("status"),
	}

	sql := `INSERT INTO papers(id, title, authors, status) VALUES(?, ?, ?, ?)`

	result, err := Db.Exec(sql, req.Id, req.Title, req.Authors, req.Status)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (pm *paperModel) UpdatePaperStatus(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	sql := `UPDATE papers SET status = ? WHERE id = ?`

	newStatus := "Reviewed" // Example status
	if r.FormValue("status") == "Reviewed" {
		newStatus = "Published"
	}

	result, err := Db.Exec(sql, newStatus, r.FormValue("id"))

	if err != nil {
		return result, err
	}

	return result, nil
}

func (pm *paperModel) DeletePaper(r *http.Request) (sql.Result, error) {
	err := r.ParseForm()

	if err != nil {
		return nil, nil
	}

	sql := `DELETE FROM papers WHERE id = ?`

	result, err := Db.Exec(sql, r.FormValue("id"))

	if err != nil {
		return result, err
	}

	return result, nil
}
