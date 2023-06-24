package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

///////////////////////// TIPO CLIENTES /////////////////////////

func getTipoClientes(w http.ResponseWriter, r *http.Request) {

	tipoClientes, err := getTipoClientesController()
	if err != nil {
		fmt.Fprintf(w, "No tipo clientes found")
	}
	if tipoClientes != nil {
		json.NewEncoder(w).Encode(tipoClientes)
	}

}

func updateTipoCliente(w http.ResponseWriter, r *http.Request) {

	var obj TipoCliente

	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateTipoClienteController(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func createTipoCliente(w http.ResponseWriter, r *http.Request) {

	var obj TipoCliente

	fmt.Println("r.Body", r.Body)

	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = createTipoClienteController(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func deleteTipoCliente(w http.ResponseWriter, r *http.Request) {

	var obj TipoCliente

	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = deleteTipoClienteController(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

///////////////////////// CLIENTES /////////////////////////

func getClientes(w http.ResponseWriter, r *http.Request) {

	clientes, err := getClientesController()
	if err != nil {
		fmt.Fprintf(w, "No clientes found")
	}
	if clientes != nil {
		json.NewEncoder(w).Encode(clientes)
	}
}

func createCliente(w http.ResponseWriter, r *http.Request) {

	var obj Cliente

	fmt.Println("r.Body", r.Body)

	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = createClienteController(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func updateCliente(w http.ResponseWriter, r *http.Request) {

	var obj Cliente

	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = updateClienteController(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func deleteCliente(w http.ResponseWriter, r *http.Request) {

	var obj Cliente

	err := json.NewDecoder(r.Body).Decode(&obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = deleteClienteController(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func main() {

	r := mux.NewRouter()

	// TIPO CLIENTES
	// get tipo clientes
	r.HandleFunc("/tipoClientes", getTipoClientes).Methods("GET")
	// create tipo cliente
	r.HandleFunc("/tipoCliente", createTipoCliente).Methods("POST")
	// update tipo cliente
	r.HandleFunc("/tipoCliente", updateTipoCliente).Methods("PUT")
	// delete tipo cliente
	r.HandleFunc("/tipoCliente", deleteTipoCliente).Methods("DELETE")

	// CLIENTES
	// get clientes
	r.HandleFunc("/clientes", getClientes).Methods("GET")
	// create cliente {"nombre":"clienteeeeeeeeeeeeeeeeeeeeeeee","dni":"12676578","telefono":"9872471","correo":"sanki2@gmail.com","estado":"A","idTipoCliente":2}
	r.HandleFunc("/cliente", createCliente).Methods("POST")
	// update cliente must include idCliente: {"idCliente":4,"nombre":"sannnnnnnnnnnnn","dni":"12676578","telefono":"9872471","correo":"sanki2@gmail.com","estado":"A","idTipoCliente":2}
	r.HandleFunc("/cliente", updateCliente).Methods("PUT")
	// delete cliente must include idCliente: {"idCliente":4}
	r.HandleFunc("/cliente", deleteCliente).Methods("DELETE")

	r.HandleFunc("/", homePage)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
