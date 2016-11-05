package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	Type        string
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

//wrkPath represents the working path
var wrkPath string

//ReadConfig reads the config file based on location passed interface{}
func ReadConfig(f string) {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(-1)
	}
	wrkPath = path.Dir(f)
	fmt.Println(wrkPath)
	configInMem, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("error reading configuration")
	}
	err = json.Unmarshal([]byte(configInMem), &wrkConfig)
	if err != nil {
		panic(err)
	}
}

//GetSchemaScripts returns the list of schema scripts from the config
func GetSchemaScripts() []string {
	var scripts []string
	if len(wrkConfig.Scripts.Schema) > 0 {
		for i := range wrkConfig.Scripts.Schema {
			scripts = append(scripts, ResolvePath(wrkConfig.Scripts.Schema[i]))
		}
		return scripts
	}
	return []string{}
}

//GetProcessScripts returns the list of process scripts from the config
func GetProcessScripts() []string {
	var scripts []string
	if len(wrkConfig.Scripts.Process) > 0 {
		for i := range wrkConfig.Scripts.Process {
			scripts = append(scripts, ResolvePath(wrkConfig.Scripts.Process[i]))
		}
		return scripts
	}
	return []string{}
}

//ResolvePath returns the fully qualified path of the name of the file passed in
func ResolvePath(p string) string {
	return wrkPath + p
}
