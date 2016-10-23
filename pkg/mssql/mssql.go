package mssql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/denisenkom/go-mssqldb" //for accessing ms sql server
)

var db *sql.DB

//OpenCn opens a connection
func OpenCn(cn string) {
	var err error
	db, err = sql.Open("mssql", cn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

//RunScripts executes the schema scripts
func RunScripts(s []string) {
	err := db.Ping()
	if err != nil {
		fmt.Println("No database connection. Cannnot run schema scripts")
		os.Exit(-1)
	}
	if len(s) > 0 {
		for _, script := range s {
			ExecScript(ReadScript(script))
		}
	} else {
		fmt.Println("No schema files processed.")
	}
}

//ExecScript executes a script
func ExecScript(s string) {
	r, err := db.Exec(s)
	if err != nil {
		panic(err)
	}
	id, err := r.LastInsertId()
	if err != nil {
	}
	af, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last inserted id:", id, "Rows affected:", af)
}

func ReadScript(s string) string {
	qry, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("Could not read file", s)
		os.Exit(-1)
	}
	return string(qry)
}
