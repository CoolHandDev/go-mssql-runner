package config

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeConnectionString(t *testing.T) {
	Convey("Given that all connection values are set", t, func() {
		cases := make(map[string]MssqlCn)
		//Case 1
		cn1 := MssqlCn{UserName: "testuser", Password: "testpassword", Server: "testhost", Database: "testdatabase", Port: "1433", AppName: "test-app-name", CnTimeout: "600"}
		cases["user id=testuser"+";password=testpassword"+";server=testhost"+";database=testdatabase"+";port=1433"+";connection timeout=600"+";app name=test-app-name"] = cn1
		//Case 2
		cn2 := MssqlCn{UserName: "anotheruser", Password: "anotherpassword", Server: "anotherhost", Database: "anotherdatabase", Port: "1433", AppName: "another-app-name", CnTimeout: "900"}
		cases["user id=anotheruser"+";password=anotherpassword"+";server=anotherhost"+";database=anotherdatabase"+";port=1433"+";connection timeout=900"+";app name=another-app-name"] = cn2
		//Case 3
		cn3 := MssqlCn{UserName: "anotheruser", Password: "", Server: "anotherhost", Database: "anotherdatabase", Port: "1433", AppName: "another-app-name", CnTimeout: "900"}
		cases["user id=anotheruser"+";password="+";server=anotherhost"+";database=anotherdatabase"+";port=1433"+";connection timeout=900"+";app name=another-app-name"] = cn3
		Convey("When they are used to construct a string", func() {

			Convey("The output should be a connection string consisting of all those values", func() {
				for k, v := range cases {
					So(k, ShouldEqual, GetCnString(v))
				}
			})

		})

	})
}

//TestReadConfig tests if the config file is being read and unmarshalled
func TestReadConfig(t *testing.T) {
	var result PrjConfig
	var expected PrjConfig

	expected.Name = "example project"
	expected.Description = "description for project"
	expected.Type = "data analysis"
	expected.Version = "1.0.0"
	expected.Scripts.Schema = []string{"/schema/schema1.sql", "/schema/schema2.sql"}
	expected.Scripts.Process = []string{"/process/process1.sql", "/process/process2.sql"}
	configInMem, err := ioutil.ReadFile("../../mssqlrun.conf.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configInMem, &result)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

//TestGetSchemaScripts tests if list of scrips is returned
func TestGetSchemaScripts(t *testing.T) {

}
