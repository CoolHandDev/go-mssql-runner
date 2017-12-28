package mssql

import (
	"database/sql"
	"errors"
	"io/ioutil"
	"time"

	_ "github.com/denisenkom/go-mssqldb" //for accessing ms sql server
	log "github.com/sirupsen/logrus"
)

//Gdb is the database object
var Gdb *sql.DB

//NewPool prepares a new connection pool
func NewPool(cn string) {
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
			log.WithFields(log.Fields{"file_name": script}).Info("Executing script file")
			_, err := ExecScript(Gdb, ReadScript(script)) //even gdb is global, pass it in so mocking is possible
			if err != nil {
				log.WithFields(log.Fields{"file_name": script}).Error(err)
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
		log.Fatal("No database connection. Cannnot run schema scripts")
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
		log.WithFields(log.Fields{"file_name": s}).Fatal(err, s)
	}
	return string(qry)
}

//queryTimer returns a timer function
func queryTimer(n string) func() {
	start := time.Now()
	return func() {
		duration := time.Now().Sub(start)
		log.WithFields(log.Fields{"file_name": n}).Info("Completed in ", duration)
	}
}
