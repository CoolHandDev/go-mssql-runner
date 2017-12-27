package mssql

import (
	"io/ioutil"
	"testing"

	log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var processFiles = []string{
	"../../example/process/process1.sql",
	"../../example/process/process2.sql",
}
var schemaFiles = []string{
	"../../example/schema/schema1.sql",
	"../../example/schema/schema2.sql",
}

func TestReadScript(t *testing.T) {
	Convey("Given that a script file is read", t, func() {

		Convey("When process1.sql is read", func() {
			processFile1 := processFiles[0]
			scpTxt := ReadScript(processFile1)
			expected := `select 'process 1'`

			Convey("The contents of process1.sql should be loaded", func() {
				So(scpTxt, ShouldEqual, expected)
			})
		})

		Convey("When schema1.sql is read", func() {
			schemaFile1 := schemaFiles[0]
			scpTxt := ReadScript(schemaFile1)
			expected := `select 'schema 1'`

			Convey("The contents of schema1 should be loaded", func() {
				So(scpTxt, ShouldEqual, expected)
			})
		})
	})
}

func TestExecScript(t *testing.T) {
	Convey("Given that a script is executed", t, func() {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		Convey("When process1.sql is executed", func() {
			processFile1 := processFiles[0]
			mock.ExpectExec(processFile1).WillReturnResult(sqlmock.NewResult(1, 1))
			_, _ = ExecScript(db, processFile1)

			Convey("The expectations should be fulfilled", func() {
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
				So(mock.ExpectationsWereMet, ShouldNotEqual, err)
			})
		})
	})
}

func TestRunScripts(t *testing.T) {
	Convey("Given that a set of scripts are executed", t, func() {
		log.SetOutput(ioutil.Discard)
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()
		Gdb = db

		Convey("When process scripts are executed", func() {
			mock.ExpectExec(`select 'process 1'`).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(`select 'process 2'`).WillReturnResult(sqlmock.NewResult(1, 1))
			i, err := RunScripts(processFiles)

			Convey("There should be no errors", func() {
				if err != nil {
					t.Errorf("errors were encountered")
				}
				So(err, ShouldBeNil)
			})
			Convey("Count of files run is 2", func() {
				So(i, ShouldEqual, 2)
			})
			Convey("The expectations should be fulfilled", func() {
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
				So(mock.ExpectationsWereMet, ShouldNotEqual, err)
			})
		})

		Convey("When schema scripts are executed", func() {
			mock.ExpectExec(`select 'schema 1'`).WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectExec(`select 'schema 2'`).WillReturnResult(sqlmock.NewResult(1, 1))
			i, err := RunScripts(schemaFiles)

			Convey("There should be no errors", func() {
				if err != nil {
					t.Errorf("errors were encountered")
				}
				So(err, ShouldBeNil)
			})
			Convey("Count of files run is 2", func() {
				So(i, ShouldEqual, 2)
			})
			Convey("The expectations should be fulfilled", func() {
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("there were unfulfilled expectations: %s", err)
				}
				So(mock.ExpectationsWereMet, ShouldNotEqual, err)
			})
		})

	})
}
