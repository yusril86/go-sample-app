package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateEmployee(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "POST":
			r.ParseForm()
			nama := r.Form["nama"][0]
			nip := r.Form["nip"][0]
			alamat := r.Form["alamat"][0]
			nohp := r.Form["nohp"][0]
			email := r.Form["email"][0]
			_, err := db.Exec("INSERT INTO employees(nama, nip, alamat,nohp,email) VALUES(?, ?, ?, ?, ?)", nama, nip, alamat, nohp, email)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)

			return

		case "GET":
			fp := filepath.Join("views", "create.html")

			tmpl, err := template.ParseFiles(fp)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
