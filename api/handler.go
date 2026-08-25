package api

import (
	"encoding/json"
	"librarytransfer/query"
	"librarytransfer/storage"
	"net/http"
)

type Handler struct{ Store *storage.Store }

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/readers" {
		http.NotFound(w, r)
		return
	}
	items, e := query.List(h.Store, query.ReaderQuery{})
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
