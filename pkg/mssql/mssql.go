package mssql

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb" //for accessing ms sql server
)

var db *sql.DB

//OpenCn opens a connection
func OpenCn(cn string) {
	var err error
	db, err = sql.Open("mssql", cn)
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
			log.Println("Executing script file", "=", script)
			ExecScript(ReadScript(script))
		}
	} else {
		log.Println("No schema files processed.")
	}
}

//ExecScript executes a script
func ExecScript(s string) {
	//r, err := db.Exec(s)
	_, err := db.Exec(s)
	if err != nil {
		log.Fatal(err)
	}
	//id, err := r.LastInsertId()
	if err != nil {
	}
	//af, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(progress.Prefix("Inserted id (" + strconv.FormatInt(id, 10) + ")" + " Rows affected (" + strconv.FormatInt(af, 10) + ")"))

}

//ReadScript loads script file
func ReadScript(s string) string {
	qry, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err, s)
	}
	return string(qry)
}
