package api

import (
	"encoding/json"
	"net/http"

	"github.com/user123-boop/fin_proj/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeError(w, "wrong method", http.StatusMethodNotAllowed)
	}
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, "Ошибка десериализации JSON", http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeError(w, "Не указан идентификатор задачи", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeError(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeError(w, "Ошибка при обновлении задачи", http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]interface{}{})
}
