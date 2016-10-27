package config

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

//TestGetCnString tests if connection string if being constructed
func TestGetCnString(t *testing.T) {
	cn := new(MssqlCn)
	cn.UserName = "testuser"
	cn.Password = "testpassword"
	cn.Server = "testhost"
	cn.Database = "testdatabase"
	cn.Port = "1433"
	cn.AppName = "test-app-name"
	cn.CnTimeout = "600"
	result := GetCnString(*cn)
	expected := "user id=testuser" +
		";password=testpassword" +
		";server=testhost" +
		";database=testdatabase" +
		";port=1433" +
		";connection timeout=600" +
		";app name=test-app-name"
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

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
