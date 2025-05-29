package api

import (
	"net/http"
	"time"

	"github.com/user123-boop/fin_proj/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		writeError(w, "wrong method", http.StatusMethodNotAllowed)
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		writeError(w, "id is empty", http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeError(w, "task not found", http.StatusNotFound)
		return
	}

	if task.Repeat == "" {
		// One-time task - delete it
		err = db.DeleteTask(id)
		if err != nil {
			writeError(w, "delete task fail", http.StatusInternalServerError)
			return
		}
	} else {
		// Recurring task - calculate next date
		now := time.Now()
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeError(w, "nextdate error", http.StatusInternalServerError)
			return
		}

		err = db.UpdateDate(id, next)
		if err != nil {
			writeError(w, "update task fail", http.StatusInternalServerError)
			return
		}
	}

	writeJSON(w, map[string]interface{}{})
}
