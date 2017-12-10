package config

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var testConfig = "../../example/mssqlrun.conf.json"

func TestMakeConnectionString(t *testing.T) {

	Convey("Given that all connection values are set", t, func() {
		cases := make(map[string]MssqlCn)
		//Case 1
		cn1 := MssqlCn{UserName: "testuser", Password: "testpassword", Server: "testhost", Database: "testdatabase", Port: "1433", AppName: "test-app-name", CnTimeout: "600", LogLevel: "4"}
		cases["user id=testuser"+";password=testpassword"+";server=testhost"+";database=testdatabase"+";port=1433"+";connection timeout=600"+";app name=test-app-name"+";log=4"] = cn1
		//Case 2
		cn2 := MssqlCn{UserName: "anotheruser", Password: "anotherpassword", Server: "anotherhost", Database: "anotherdatabase", Port: "1433", AppName: "another-app-name", CnTimeout: "900", LogLevel: "63"}
		cases["user id=anotheruser"+";password=anotherpassword"+";server=anotherhost"+";database=anotherdatabase"+";port=1433"+";connection timeout=900"+";app name=another-app-name"+";log=63"] = cn2
		//Case 3
		cn3 := MssqlCn{UserName: "anotheruser", Password: "", Server: "anotherhost", Database: "anotherdatabase", Port: "1433", AppName: "another-app-name", CnTimeout: "900", LogLevel: ""}
		cases["user id=anotheruser"+";password="+";server=anotherhost"+";database=anotherdatabase"+";port=1433"+";connection timeout=900"+";app name=another-app-name"+";log="] = cn3

		Convey("When they are used to construct a connectrion string", func() {

			Convey("The output should be a MSSQL connection string consisting of all those values", func() {
				for k, v := range cases {
					So(GetCnString(v), ShouldEqual, k)
				}
			})

		})

	})
}

func TestLoadConfiguration(t *testing.T) {

	Convey("Given that a configuration file is specified", t, func() {
		var expected PrjConfig

		expected.Name = "example project"
		expected.Description = "description for project"
		expected.Type = "data analysis"
		expected.Version = "1.0.0"
		expected.Scripts.Schema = []string{"/schema/schema1.sql", "/schema/schema2.sql"}
		expected.Scripts.Process = []string{"/process/process1.sql", "/process/process2.sql"}

		Convey("When it is loaded", func() {

			ReadConfig(testConfig)
			result := wrkConfig
			Convey("The values from the file should be available to the app", func() {
				So(result, ShouldResemble, expected)
			})

			Convey("The schema array should contain schema file names", func() {
				Convey("The first schema file should exist", func() {
					So(result.Scripts.Schema[0], ShouldEqual, expected.Scripts.Schema[0])
				})
				Convey("The last schema file should exist", func() {
					resultIdx := len(result.Scripts.Schema) - 1
					expectedIdx := len(expected.Scripts.Schema) - 1
					So(result.Scripts.Schema[resultIdx], ShouldEqual, expected.Scripts.Schema[expectedIdx])
				})
			})

			Convey("The process array should contain process file names", func() {
				Convey("The first process file should exist", func() {
					So(result.Scripts.Process[0], ShouldEqual, expected.Scripts.Process[0])
				})
				Convey("The last process file should exist", func() {
					resultIdx := len(result.Scripts.Process) - 1
					expectedIdx := len(expected.Scripts.Process) - 1
					So(result.Scripts.Process[resultIdx], ShouldEqual, expected.Scripts.Process[expectedIdx])
				})
			})

			Convey("The config should return a list of scripts", func() {
				Convey("A list of schema scripts should be returned", func() {
					schemaList := GetSchemaScripts()
					So(schemaList, ShouldNotBeEmpty)
				})
				Convey("A list of process scripts should be returned", func() {
					processList := GetProcessScripts()
					So(processList, ShouldNotBeEmpty)
				})
			})

			Convey("The list of scripts should be able to resolve to files", func() {
				Convey("A schema script file should be able to be accessed", func() {
					schemaList := GetSchemaScripts()
					So(schemaList[0], shouldFileExist)
				})
				Convey("A process script file should be able to be accessed", func() {
					processList := GetProcessScripts()
					So(processList[0], shouldFileExist)
				})
			})

		})
	})
}

//Custom assertion to check existence of file
func shouldFileExist(actual interface{}, expected ...interface{}) string {
	_, err := os.Stat(actual.(string))
	if os.IsNotExist(err) {
		return "Expected file to exist, but it did not"
	}
	return ""
}
