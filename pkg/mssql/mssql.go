package mssql

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	"time"

	_ "github.com/denisenkom/go-mssqldb" //for accessing ms sql server
)

var db *sql.DB

//OpenCn opens a connection
func OpenCn(cn string) {
	var err error
	db, err = sql.Open("sqlserver", cn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

//RunScripts executes the schema scripts
func RunScripts(s []string) {
	err := db.Ping()
	if err != nil {
		log.Println("No database connection. Cannnot run schema scripts")
		os.Exit(-1)
	}
	if len(s) > 0 {
		for _, script := range s {
			log.Println("-------------------------------")
			timer := queryTimer(script)
			log.Println("Executing script file", "=", script)
			ExecScript(ReadScript(script))
			timer()
			log.Println("-------------------------------")
		}
	} else {
		log.Println("No schema files processed.")
	}
}

//ExecScript executes a script
func ExecScript(s string) {
	_, err := db.Exec(s)
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
