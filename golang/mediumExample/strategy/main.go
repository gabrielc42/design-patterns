// strategy:
// set of behaviors ->
// interchangeable objects inside original context object
// don't need certain format to run

// connect multiple types of DB
// able to implement new type of DB connection

package main

import "fmt"

type IDBconnection interface {
	Connect()
}

type SqlConnection struct {
	connectionString string
}

func (con SqlConnection) Connect() {
	fmt.Println("Sql " + con.connectionString)
}

type OracleConnection struct {
	connectionString string
}

func (con OracleConnection) Connect() {
	fmt.Println("Oracle " + con.connectionString)
}

type DBConnection struct {
	db IDBconnection
}

func (con DBConnection) DBConnect() {
	con.db.Connect()
}

func main() {
	sqlConnection := SqlConnection{connectionString: "Connection is connected :) "}
	con := DBConnection{db: sqlConnection}
	con.DBConnect()

	orcConnection := OracleConnection{connectionString: "Connection is connected :P "}
	con2 := DBConnection{db: orcConnection}
	con2.DBConnect()
}
