package api

import (
	"net/http"

	"github.com/user123-boop/fin_proj/pkg/db"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "wrong method", http.StatusMethodNotAllowed)
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	writeJSON(w, task)
}
