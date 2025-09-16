package controller

import (
	"database/sql"
	"net/http"
)

func DeleteEmployee(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")

		_, err := db.Exec("DELETE FROM employees WHERE id=?", id)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employee", http.StatusMovedPermanently)

	}
}
