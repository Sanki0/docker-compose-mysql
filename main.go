package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func connectionDB() *sql.DB {
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_database")
	if err != nil {
		panic(err.Error())
	}
	return db
}
func chkError(err error) {
	if err != nil {
		panic(err)
	}
}

func PingDb(db *sql.DB) {
	err := db.Ping()
	chkError(err)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getTipoClientes() ([]TipoCliente, error) {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "SELECT * FROM tipoCliente"
	rows, err := db.Query(statement)
	chkError(err)
	defer rows.Close()

	tipoClientes := []TipoCliente{}

	for rows.Next() {
		var t TipoCliente
		if err := rows.Scan(&t.IdTipoCliente, &t.Nombre, &t.Detalle); err != nil {
			return nil, err
		}
		tipoClientes = append(tipoClientes, t)
	}

	return tipoClientes, nil
}

func getTipoClientesController(w http.ResponseWriter, r *http.Request) {

	tipoClientes, err := getTipoClientes()
	if err != nil {
		fmt.Fprintf(w, "No tipo clientes found")
	}
	if tipoClientes != nil {
		json.NewEncoder(w).Encode(tipoClientes)
	}

}

func getClientes() ([]Cliente, error) {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "SELECT * FROM cliente"
	rows, err := db.Query(statement)
	chkError(err)

	clientes := []Cliente{}

	for rows.Next() {
		var c Cliente
		if err := rows.Scan(&c.IdCliente, &c.Nombre, &c.Dni, &c.Telefono, &c.Correo, &c.Estado, &c.IdTipoCliente); err != nil {
			return nil, err
		}
		clientes = append(clientes, c)
	}

	return clientes, nil
}

func getClientesController(w http.ResponseWriter, r *http.Request) {

	clientes, err := getClientes()
	if err != nil {
		fmt.Fprintf(w, "No clientes found")
	}
	if clientes != nil {
		json.NewEncoder(w).Encode(clientes)
	}
}

func main() {

	r := mux.NewRouter()

	// get tipo clientes
	r.HandleFunc("/tipoClientes", getTipoClientesController).Methods("GET")

	// get clientes
	r.HandleFunc("/clientes", getClientesController).Methods("GET")

	r.HandleFunc("/", homePage)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
