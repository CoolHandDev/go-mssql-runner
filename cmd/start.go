// Copyright © 2016 The Go MSSQL Runner Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/coolhanddev/go-mssql-runner/pkg/config"
	"github.com/coolhanddev/go-mssql-runner/pkg/mssql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configFile string
var userName string
var password string
var server string
var database string
var port string
var cnTimeout string
var appName string
var logLevel string
var cn config.MssqlCn
var logToFile string
var logFileName os.File
var logFormat string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start running scripts",
	Long: `

The start command kicks off the execution of the scripts listed in the 
configuration json file specified in the --conf flag. Connection
information to MS SQL Server must be passed in.

The minimum required database connection information are: username,
password, server and database.

The minimum required values to run the command can be set in environment
variables.

GOSQLR_CONFIGFILE
GOSQLR_USERNAME
GOSQLR_PASSWORD
GOSQLR_SERVER
GOSQLR_DATABASE

Specifying parameter values will override the environment settings. 

Examples:
Passing parameters
> go-mssql-runner start -u dbuser -p secretpass -s 172.0.0.1 -d mydatabasename -c x:\workspace\mssqlrun.conf.json

When environment variables are set, then just run  
> go-mssql-runner start

Different log levels can be set via the -l flag.

0 no logging
1 log errors
2 log messages 
4 log rows affected 
8 trace sql statements 
16 log statement parameters 
32 log transaction begin/end 
63 full logging

`,
	Run: start,
}

func start(cmd *cobra.Command, args []string) {
	if server == "" || database == "" || configFile == "" {
		//if required parameters are not met, check environment variables
		if os.Getenv("GOSQLR_CONFIGFILE") == "" || os.Getenv("GOSQLR_SERVER") == "" || os.Getenv("GOSQLR_DATABASE") == "" {
			fmt.Println("Please pass in the required values. ")
			fmt.Println("go-mssql-runner start -h for more information.")
			os.Exit(1)
		} else if os.Getenv("GOSQLR_CONFIGFILE") != "" && os.Getenv("GOSQLR_USERNAME") != "" &&
			os.Getenv("GOSQLR_PASSWORD") != "" && os.Getenv("GOSQLR_SERVER") != "" &&
			os.Getenv("GOSQLR_DATABASE") != "" {
			configFile = os.Getenv("GOSQLR_CONFIGFILE")
			userName = os.Getenv("GOSQLR_USERNAME")
			password = os.Getenv("GOSQLR_PASSWORD")
			server = os.Getenv("GOSQLR_SERVER")
			database = os.Getenv("GOSQLR_DATABASE")
		}
	}
	cn.UserName = userName
	cn.Password = password
	cn.Server = server
	cn.Database = database
	cn.Port = port
	cn.AppName = appName
	cn.CnTimeout = cnTimeout
	cn.LogLevel = logLevel
	startTime := time.Now()
	//set up logging. we want to log both to stdout and to a file
	if logFormat == "JSON" {
		log.SetFormatter(&log.JSONFormatter{TimestampFormat: "01-02-2006 15:04:05"})
	} else {
		log.SetFormatter(&log.TextFormatter{TimestampFormat: "01-02-2006 15:04:05", FullTimestamp: true})
	}
	if logToFile != "" {
		//if file already exists then append. log rotation done manually by user
		logFileName, err := os.OpenFile(logToFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}
		mw := io.MultiWriter(os.Stdout, logFileName)
		log.SetOutput(mw)
	}
	newDbPool(cn)
	log.WithFields(log.Fields{"config_file": configFile}).Info("loading configuration")
	config.ReadConfig(configFile)
	log.Info("================================================")
	log.Info("Executing schema scripts")
	_, err := mssql.RunScripts(config.GetSchemaScripts())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("================================================")
	log.Info("Executing process scripts")
	_, err = mssql.RunScripts(config.GetProcessScripts())
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(startTime)
	log.WithFields(log.Fields{"log_file": logToFile}).Info("Log file created")
	log.Info("Total time elapsed", "=", elapsed)
	logFileName.Close()
}

func newDbPool(c config.MssqlCn) {
	log.WithFields(log.Fields{"server": cn.Server, "database": cn.Database}).Info("opening database")
	mssql.NewPool(config.GetCnString(c))
}

func init() {
	RootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	logLevelMsg := "Specifies level of verbosity for SQL log output. See start command help (above) for details"
	startCmd.Flags().StringVarP(&configFile, "conf", "c", "", "The configuration file")
	startCmd.Flags().StringVarP(&userName, "username", "u", "", "SQL Server user name. *Required")
	startCmd.Flags().StringVarP(&password, "password", "p", "", "SQL Server user password. *Required")
	startCmd.Flags().StringVarP(&server, "server", "s", "", "SQL Server host. *Required")
	startCmd.Flags().StringVarP(&database, "database", "d", "", "Database to work on. *Required")
	startCmd.Flags().StringVarP(&port, "port", "", "1433", "Host port number")
	startCmd.Flags().StringVarP(&cnTimeout, "timeout", "t", "14400", "Connection timeout in seconds")
	startCmd.Flags().StringVarP(&appName, "appname", "a", "go-mssql-runner", "App name to show in db calls. Useful for SQL Profiler")
	startCmd.Flags().StringVarP(&logLevel, "loglevel", "l", "0", logLevelMsg)
	startCmd.Flags().StringVarP(&logToFile, "logfile", "", "", "File to write log to")
	startCmd.Flags().StringVarP(&logFormat, "logformat", "", "text", "Format of log: JSON or text")
}
