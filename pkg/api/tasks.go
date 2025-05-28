package api

import (
	"net/http"
	"strconv"

	"github.com/user123-boop/fin_proj/pkg/db"
)

const limitConst = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = limitConst
	}
	tasks, err := db.Tasks(limit)
	if err != nil {
		writeError(w, "Ошибка при получении задач", http.StatusInternalServerError)
		return
	}
	writeJSON(w, TasksResp{Tasks: tasks})
}
