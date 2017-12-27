package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
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
	LogLevel  string
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

var wrkConfig PrjConfig

//wrkPath represents the working path
var wrkPath string

//GetCnString returns a sql server connection string
func GetCnString(c MssqlCn) string {
	return "user id=" + c.UserName +
		";password=" + c.Password +
		";server=" + c.Server +
		";database=" + c.Database +
		";port=" + c.Port +
		";connection timeout=" + c.CnTimeout +
		";app name=" + c.AppName +
		";log=" + c.LogLevel
}

//ReadConfig reads the config file based on location passed interface{}
func ReadConfig(f string) {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		log.WithFields(log.Fields{"file_name": f}).Fatal(err)
	}
	wrkPath = filepath.Dir(f)
	configInMem, err := ioutil.ReadFile(f)
	if err != nil {
		log.WithFields(log.Fields{"file_name": f}).Fatal(err)
	}
	err = json.Unmarshal([]byte(configInMem), &wrkConfig)
	if err != nil {
		log.WithFields(log.Fields{"file_name": f}).Fatal(err)
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
