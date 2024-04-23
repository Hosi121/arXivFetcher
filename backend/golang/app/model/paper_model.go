package model

import (
	"database/sql"
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/oklog/ulid"
)

type PaperModel interface {
	FetchPaper(r *http.Request) ([]*Paper, error)
	AnalyzeCitations(r *http.Request) ([]Citation, error)
	ExpandGraph(r *http.Request) (*Graph, error)
	DeletePaper(r *http.Request) (sql.Result, error)
}

type paperModel struct {
	Db *sql.DB
}

type Paper struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Authors string `json:"authors"`
	Status  string `json:"status"`
}

type Node struct {
	ID    string
	Value interface{}
}
type Edge struct {
	Start  *Node
	End    *Node
	Weight float64
}

type Citation struct {
	// ここに引用情報のフィールドを定義
}

type Graph struct {
	Nodes map[string]*Node
	Edges []*Edge
	// ここにグラフ構造のフィールドを定義
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make([]*Edge, 0),
	}
}

func (g *Graph) AddNode(node *Node) error {
	if _, exists := g.Nodes[node.ID]; exists {
		return errors.New("Node already exists with the same ID")
	}
	g.Nodes[node.ID] = node
	return nil
}
func (g *Graph) AddEdge(startID, endID string, weight float64) error {
	startNode, ok := g.Nodes[startID]
	if !ok {
		return errors.New("start node does not exist")
	}
	endNode, ok := g.Nodes[endID]
	if !ok {
		return errors.New("snd node does not exist")
	}
	edge := &Edge{
		Start:  startNode,
		End:    endNode,
		Weight: weight,
	}
	g.Edges = append(g.Edges, edge)
	return nil
}
func CreatePaperModel(db *sql.DB) PaperModel {
	return &paperModel{Db: db}
}

func (pm *paperModel) FetchPaper(r *http.Request) ([]*Paper, error) {
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
	return []Citation{}, nil // Todo
}

func (pm *paperModel) ExpandGraph(r *http.Request) (*Graph, error) {
	return &Graph{}, nil // Todo
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
