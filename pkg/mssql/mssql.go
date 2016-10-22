package mssql

import (
	"database/sql"
	"fmt"
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

func CreateSchema(s []string) {
	err := db.Ping()
	if err != nil {
		fmt.Println("No database connection. Cannnot run schema scripts")
		os.Exit(-1)
	}
	if len(s) > 0 {
		for _, schema := range s {
			fmt.Println(schema)
		}
	} else {
		fmt.Println("No schema files processed.")
	}
}

func ExecScripts(s []string) {

}
