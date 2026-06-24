package handlers

import (
	"io"
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHander(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func UploadHander(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "cannot read file: " + err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "cannot read content", http.StatusInternalServerError)
		return
	}

	result := service.Converter(string(data))

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}