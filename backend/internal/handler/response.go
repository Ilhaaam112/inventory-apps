package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Helper bersama untuk handler transaksi.

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func idFromPath(r *http.Request, prefix string) (int, error) {
	return strconv.Atoi(strings.Trim(strings.TrimPrefix(r.URL.Path, prefix), "/"))
}

func queryInt(r *http.Request, key string) int {
	v, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		return 0
	}
	return v
}
