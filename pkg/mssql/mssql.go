package mssql

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	"time"

	_ "github.com/denisenkom/go-mssqldb" //for accessing ms sql server
)

var gdb *sql.DB

//OpenCn opens a connection
func OpenCn(cn string) {
	var err error
	gdb, err = sql.Open("sqlserver", cn)
	if err != nil {
		log.Fatal(err)
	}
	err = gdb.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

//RunScripts executes the schema scripts
func RunScripts(s []string) {
	if len(s) > 0 {
		for _, script := range s {
			log.Println("-------------------------------")
			timer := queryTimer(script)
			log.Println("Executing script file", "=", script)
			ExecScript(gdb, ReadScript(script)) //gdb is global, but we need to be able to mock db in testing
			timer()
			log.Println("-------------------------------")
		}
	} else {
		log.Println("No schema files processed.")
	}
}

//ExecScript executes a script
func ExecScript(db *sql.DB, s string) {
	err := db.Ping()
	if err != nil {
		log.Println("No database connection. Cannnot run schema scripts")
		os.Exit(-1)
	}
	_, err = db.Exec(s)
	if err != nil {
		log.Fatal(err)
	}
}

//ReadScript loads script file
func ReadScript(s string) string {
	qry, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err, s)
	}
	return string(qry)
}

//queryTimer returns a timer function
func queryTimer(n string) func() {
	start := time.Now()
	return func() {
		duration := time.Now().Sub(start)
		log.Println(n, "completed in", duration)
	}
}
