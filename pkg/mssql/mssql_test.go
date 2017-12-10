package mssql

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var processFile1 = "../../example/process/process1.sql"
var schemaFile1 = "../../example/schema/schema1.sql"

func TestReadScript(t *testing.T) {
	Convey("Given that a script file is read", t, func() {

		Convey("When process1.sql is read", func() {
			scpTxt := ReadScript(processFile1)
			expected := `select 'process 1'`

			Convey("The contents of process1.sql should be loaded", func() {
				So(scpTxt, ShouldEqual, expected)
			})
		})

		Convey("When schema1.sql is read", func() {
			scpTxt := ReadScript(schemaFile1)
			expected := `select 'schema 1'`

			Convey("The contents of schema1 should be loaded", func() {
				So(scpTxt, ShouldEqual, expected)
			})
		})
	})
}
