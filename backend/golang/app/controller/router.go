package controller

import (
    "net/http"
    "os"
)

type Router interface {
    FetchPaper(w http.ResponseWriter, r *http.Request)
    AnalyzeCitations(w http.ResponseWriter, r *http.Request)
    ExpandGraph(w http.ResponseWriter, r *http.Request)
}

type router struct {
    pc PaperController
}

func CreateRouter(pc PaperController) Router {
    return &router{pc}
}

func setCommonHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Headers", "*")
    w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
    w.Header().Set("Content-Type", "application/json")
}

func (ro *router) FetchPaper(w http.ResponseWriter, r *http.Request) {
    if r.Method == "OPTIONS" {
        return
    }
    setCommonHeaders(w)
    ro.pc.FetchPaper(w, r)
}

func (ro *router) AnalyzeCitations(w http.ResponseWriter, r *http.Request) {
    if r.Method == "OPTIONS" {
        return
    }
    setCommonHeaders(w)
    ro.pc.AnalyzeCitations(w, r)
}

func (ro *router) ExpandGraph(w http.ResponseWriter, r *http.Request) {
    if r.Method == "OPTIONS" {
        return
    }
    setCommonHeaders(w)
    ro.pc.ExpandGraph(w, r)
}
