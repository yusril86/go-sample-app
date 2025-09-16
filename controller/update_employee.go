package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func UpdateEmployee(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "POST":
			id := r.URL.Query().Get("id")
			r.ParseForm()
			nama := r.Form["nama"][0]
			nip := r.Form["nip"][0]
			alamat := r.Form["alamat"][0]
			nohp := r.Form["nohp"][0]
			email := r.Form["email"][0]
			_, err := db.Exec("UPDATE employees SET nama=?, nip=?, alamat=?, nohp=? ,email=? WHERE id=? ", nama, nip, alamat, nohp, email, id)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)

			return

		case "GET":

			id := r.URL.Query().Get("id")
			row := db.QueryRow("SELECT nama, nip, alamat, nohp, email FROM employees WHERE id = ?", id)
			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var employee Employee
			err := row.Scan(&employee.Nama, &employee.Nip, &employee.Alamat, &employee.Nohp, &employee.Email)
			employee.Id = id
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update.html")

			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["employee"] = employee

			err = tmpl.Execute(w, data)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
