package config

import (
	"testing"
)

func TestGetCnString(t *testing.T) {
	result := GetCnString("testuser", "testpassword", "testhost", "testdatabase")
	expected := "user id=testuser;password=testpassword;server=testhost;database=testdatabase"

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}

}
