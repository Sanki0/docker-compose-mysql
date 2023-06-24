package main

import "database/sql"

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

func getTipoClientesController() ([]TipoCliente, error) {

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

func updateTipoClienteController(t TipoCliente) error {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "UPDATE tipoCliente SET nombre=?, detalle=? WHERE idTipoCliente=?"
	stmt, err := db.Prepare(statement)
	chkError(err)

	_, err = stmt.Exec(t.Nombre, t.Detalle, t.IdTipoCliente)
	chkError(err)

	return nil
}

func deleteTipoClienteController(t TipoCliente) error {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "DELETE FROM tipoCliente WHERE idTipoCliente=?"
	stmt, err := db.Prepare(statement)
	chkError(err)

	_, err = stmt.Exec(t.IdTipoCliente)
	chkError(err)

	return nil
}

func createTipoClienteController(t TipoCliente) error {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "INSERT INTO tipoCliente (nombre, detalle) VALUES (?, ?)"
	stmt, err := db.Prepare(statement)
	chkError(err)

	_, err = stmt.Exec(t.Nombre, t.Detalle)
	chkError(err)

	return nil
}

func getClientesController() ([]Cliente, error) {

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

func updateClienteController(c Cliente) error {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "UPDATE cliente SET nombre=?, dni=?, telefono=?, correo=?, estado=?, idTipoCliente=? WHERE idCliente=?"
	stmt, err := db.Prepare(statement)
	chkError(err)

	_, err = stmt.Exec(c.Nombre, c.Dni, c.Telefono, c.Correo, c.Estado, c.IdTipoCliente, c.IdCliente)
	chkError(err)

	return nil
}

func deleteClienteController(c Cliente) error {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "DELETE FROM cliente WHERE idCliente=?"
	stmt, err := db.Prepare(statement)
	chkError(err)

	_, err = stmt.Exec(c.IdCliente)
	chkError(err)

	return nil
}

func createClienteController(c Cliente) error {

	db := connectionDB()
	defer db.Close()
	PingDb(db)

	statement := "INSERT INTO cliente (nombre, dni, telefono, correo, estado, idTipoCliente) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(statement)
	chkError(err)

	_, err = stmt.Exec(c.Nombre, c.Dni, c.Telefono, c.Correo, c.Estado, c.IdTipoCliente)
	chkError(err)

	return nil
}
