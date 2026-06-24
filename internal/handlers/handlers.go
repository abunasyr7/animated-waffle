package handlers

import (
	"net/http"
)

func IndexHander(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

