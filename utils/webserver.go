package utils

import (
	"io"
	"net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Healthy\n")
}