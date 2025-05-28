package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/user123-boop/fin_proj/pkg/db"
)

func afterNow(date, now time.Time) bool {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return date.After(now)
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

/*func writeJSON(w http.ResponseWriter, data any) {
	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Ошибка маршализации JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//json.NewEncoder(w).Encode(data)
	w.Write(resp)
}*/

func writeError(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	writeJSON(w, map[string]string{"error": error})
	//w.WriteHeader(code)
}

func checkDate(task *db.Task) error {
	now := time.Now()
	today := now.Format("20060102")

	if task.Date == "" {
		task.Date = today
	}

	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("data is invalid")
	}

	if !afterNow(t, now) {
		if task.Repeat == "" {
			task.Date = today
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return fmt.Errorf("repeat is invalid")
			}
			task.Date = next
		}
	}

	return nil
}
