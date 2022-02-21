package http_handler

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Header
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write([]byte("/"))
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.Header
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Started!"))

}
