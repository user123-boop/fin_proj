package api

import (
	"net/http"
	"strconv"

	"github.com/user123-boop/fin_proj/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		writeError(w, "wrong method", http.StatusMethodNotAllowed)
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "Не указан идентификатор", http.StatusBadRequest)
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		writeError(w, "id должно быть числом", http.StatusBadRequest)
		return
	}

	err = db.DeleteTask(id)
	if err != nil {
		writeError(w, "Ошибка при удалении задачи", http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]interface{}{})
}
