package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/user123-boop/fin_proj/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "wrong method", http.StatusMethodNotAllowed)
	}
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeError(w, "decode fail", http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeError(w, "title is empty", http.StatusBadRequest)
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeError(w, "addtask fail", http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{"id": fmt.Sprint(id)})
}
