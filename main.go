package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type healthInfo struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

var startTime = time.Now()

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/healthz", healthStatus)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func healthStatus(w http.ResponseWriter, r *http.Request) {
	healthInfo := healthInfo{}
	healthInfo.Service = "Hello World Service"
	healthInfo.Status = "OK"
	healthInfo.Version = "0.0.1"
	healthInfo.Uptime = fmt.Sprint("up since ", startTime.Format("2006-01-02 15:04:05"))

	js, err := json.MarshalIndent(healthInfo,"","  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
