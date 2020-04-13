package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/gungorugur/healthcheck/internal/cache"
	"github.com/gungorugur/healthcheck/internal/database"
)

func main() {
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/badhealthcheck", badHealthcheck)
	http.ListenAndServe(":8080", nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	commit := os.Getenv("COMMIT")
	hostname, _ := os.Hostname()
	runtime := runtime.Version()
	cacheHealthy := cache.IsHealthy()
	dbHealthy := database.IsHealthy()
	alive := cacheHealthy && dbHealthy

	response := map[string]string{
		"version":      version,
		"commit":       commit,
		"hostname":     hostname,
		"runtime":      runtime,
		"cacheHealthy": strconv.FormatBool(cacheHealthy),
		"dbHealthy":    strconv.FormatBool(dbHealthy),
		"alive":        strconv.FormatBool(alive),
	}

	w.Header().Set("Content-Type", "application/json")

	if alive {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}

	json.NewEncoder(w).Encode(response)
}

func badHealthcheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"alive": "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
