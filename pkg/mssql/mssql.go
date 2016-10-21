package mssql

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb" //for accessing ms sql server
)

var db *sql.DB

//OpenCn opens a connection
func OpenCn(cn string) {
	db, err := sql.Open("mssql", cn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("was able to connect")
}
