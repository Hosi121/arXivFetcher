package controller

import (
	"crud/model" // モデル層を論文処理用に変更
	"encoding/json"
	"fmt"
	"net/http"
)

type PaperController interface {
	FetchPaper(w http.ResponseWriter, r *http.Request)
	AnalyzeCitations(w http.ResponseWriter, r *http.Request)
	ExpandGraph(w http.ResponseWriter, r *http.Request)
	DeletePaper(w http.ResponseWriter, r *http.Request)
}

type paperController struct {
	pm model.PaperModel // 論文処理のモデル
}

func CreatePaperController(pm model.PaperModel) PaperController {
	return &paperController{pm}
}

func (pc *paperController) FetchPaper(w http.ResponseWriter, r *http.Request) {
	paper, err := pc.pm.FetchPaper(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(paper)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (pc *paperController) AnalyzeCitations(w http.ResponseWriter, r *http.Request) {
	citations, err := pc.pm.AnalyzeCitations(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(citations)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (pc *paperController) ExpandGraph(w http.ResponseWriter, r *http.Request) {
	graph, err := pc.pm.ExpandGraph(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(graph)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

func (pc *paperController) DeletePaper(w http.ResponseWriter, r *http.Request) {
	result, err := pc.pm.DeletePaper(r)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	json, err := json.Marshal(result)

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
