package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// リクエストデータの構造体
type RequestData struct {
    Message string `json:"message"`
}

// レスポンスデータの構造体
type ResponseData struct {
    Response string `json:"response"`
}

// ハンドラ関数
func postDataHandler(w http.ResponseWriter, r *http.Request) {
    var requestData RequestData
    json.NewDecoder(r.Body).Decode(&requestData)
    response := ResponseData{Response: "Received: " + requestData.Message}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/data", postDataHandler).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", r))
}

