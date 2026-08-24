package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Helper bersama untuk seluruh handler.

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// writeError tidak pernah meneruskan error database atau stack trace.
// Pemanggil bertanggung jawab mengirim pesan yang aman dibaca user.
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

func queryStr(r *http.Request, key string) string {
	return strings.TrimSpace(r.URL.Query().Get(key))
}

// decodeJSON membatasi ukuran body dan menolak field yang tidak dikenal,
// supaya payload aneh dari client langsung ditolak di pintu masuk.
func decodeJSON(r *http.Request, target interface{}) error {
	r.Body = http.MaxBytesReader(nil, r.Body, 1<<20) // 1 MB
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(target)
}
