package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//MssqlCn represents the connection string
type MssqlCn struct {
	UserName  string
	Password  string
	Server    string
	Database  string
	Port      string
	CnTimeout string
	AppName   string
}

//PrjConfig represents the configuration from the json file
type PrjConfig struct {
	Name        string
	Description string
	PrjType     string
	Version     string
	Scripts     CfgScripts
}

//CfgScripts represents the list of schema and process scripts
type CfgScripts struct {
	Schema  []string
	Process []string
}

//GetCnString returns a sql server connection string
func GetCnString(c MssqlCn) string {
	return "user id=" + c.UserName +
		";password=" + c.Password +
		";server=" + c.Server +
		";database=" + c.Database +
		";port=" + c.Port +
		";connection timeout=" + c.CnTimeout +
		";app name=" + c.AppName

}

var wrkConfig PrjConfig

//ReadConfig reads the config file based on location passed in
func ReadConfig(f string) {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(-1)
	}

	configInMem, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("error reading configuration")
	}

	err = json.Unmarshal([]byte(configInMem), &wrkConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(wrkConfig)

}
