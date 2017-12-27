package mssql

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb" //for accessing ms sql server
	log "github.com/sirupsen/logrus"
)

//Gdb is the database object
var Gdb *sql.DB

func init() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "01-02-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

//OpenCn opens a connection
func OpenCn(cn string) {
	var err error
	Gdb, err = sql.Open("sqlserver", cn)
	if err != nil {
		log.Fatal(err)
	}
	err = Gdb.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

//RunScripts executes the schema scripts
func RunScripts(s []string) (c int, err error) {
	if len(s) > 0 {
		for i, script := range s {
			log.Println("-------------------------------")
			timer := queryTimer(script)
			log.Println("Executing script file", "=", script)
			_, err := ExecScript(Gdb, ReadScript(script)) //gdb is global, but we need to be able to mock db in testing
			if err != nil {
				log.Println("An error was encountered in the script:", script)
				return i, err
			}
			timer()
			log.Println("-------------------------------")
		}
		return len(s), nil
	}
	return 0, errors.New("no schema files processed")
}

//ExecScript executes a script
func ExecScript(db *sql.DB, s string) (r sql.Result, err error) {
	err = db.Ping()
	if err != nil {
		log.Println("No database connection. Cannnot run schema scripts")
		os.Exit(-1)
	}
	r, err = db.Exec(s)
	if err != nil {
		return nil, err
	}
	return r, nil
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
