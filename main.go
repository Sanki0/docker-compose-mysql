package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID               int    `json:"id"`
	Nombre           string `json:"nombre"`
	Dni              string `json:"dni"`
	Direccion        string `json:"direccion"`
	Fecha_nacimiento string `json:"fecha_nacimiento"`
}

func main() {
	db, err := sql.Open("mysql", "test_user:secret@tcp(172.18.0.2:3306)/test_database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nombre, dni, direccion, fecha_nacimiento FROM students")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var nombre, dni, direccion, fechaNacimiento string
			err := rows.Scan(&id, &nombre, &dni, &direccion, &fechaNacimiento)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "ID: %d, Nombre: %s, DNI: %s, Dirección: %s, Fecha de Nacimiento: %s\n", id, nombre, dni, direccion, fechaNacimiento)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
