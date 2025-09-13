package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Employee struct {
	Id                             string
	Nama, Nip, Alamat, Nohp, Email string
}

func IndexEmployee(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT id, nama, nip, alamat, nohp, email FROM employees")

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		var employees []Employee

		for rows.Next() {
			var employee Employee

			rows.Scan(
				&employee.Id,
				&employee.Nama,
				&employee.Nip,
				&employee.Alamat,
				&employee.Nohp,
				&employee.Email)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return

			}

			employees = append(employees, employee)

		}

		fp := filepath.Join("views", "index.html")

		tmpl, err := template.ParseFiles(fp)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["employees"] = employees

		err = tmpl.Execute(w, data)

		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
