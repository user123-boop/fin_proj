package api

import (
	"net/http"

	"github.com/user123-boop/fin_proj/pkg/db"
)

const limitConst = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := db.Tasks(limitConst)
	if err != nil {
		writeError(w, "Ошибка при получении задач", http.StatusInternalServerError)
		return
	}
	writeJSON(w, TasksResp{Tasks: tasks})
}
