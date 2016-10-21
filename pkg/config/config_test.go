package config

import (
	"testing"
)

func TestGetCnString(t *testing.T) {
	cn := new(MssqlCn)
	cn.UserName = "testuser"
	cn.Password = "testpassword"
	cn.Server = "testhost"
	cn.Database = "testdatabase"
	cn.Port = "1433"
	cn.Encrypt = "false"
	cn.AppName = "test-app-name"
	cn.CnTimeout = "600"
	result := GetCnString(*cn)
	expected := "user id=testuser" +
		";password=testpassword" +
		";server=testhost" +
		";database=testdatabase" +
		";port=1433" +
		";connection timeout=600" +
		";encrypt=false" +
		";app name=test-app-name"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

}
