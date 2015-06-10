package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type Status struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func getStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rec := Status{Message: "Ok", Timestamp: time.Now().Unix()}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if json.NewEncoder(w).Encode(rec) != nil {
		w.WriteHeader(500)
	}
}
