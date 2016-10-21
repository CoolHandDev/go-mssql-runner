package config

import (
	"testing"
)

func TestGetCnString(t *testing.T) {
	result := GetCnString("testuser", "testpassword", "testhost", "testdatabase")
	expected := "user id=testuser;password=testpassword;server=testhost;database=testdatabase"

	if result != expected {
		t.Log("test failed", "got", result, "expected", expected)
		t.Fail()
	}

}
