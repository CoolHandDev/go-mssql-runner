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

//TODO:  Allow checking of environment variables for required connection values
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/coolhanddev/go-mssql-runner/pkg/config"
	"github.com/coolhanddev/go-mssql-runner/pkg/message/progress"
	"github.com/coolhanddev/go-mssql-runner/pkg/mssql"
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
var cn config.MssqlCn

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start running scripts",
	Long: `

The start command kicks off the execution of the scripts listed in 
the configuration json file specified in the --conf flag. Connection
information to MS SQL Server must be passed in.

The minimum required connection information are: username, password,
server and database.

`,
	Run: func(cmd *cobra.Command, args []string) {
		//Do not continue if required connection parameters are not met
		if userName == "" || password == "" || server == "" || database == "" {
			fmt.Println("Please pass in the required connection values. ")
			fmt.Println("go-mssql-runner start -h for more information.")
			os.Exit(-1)
		}
		cn.UserName = userName
		cn.Password = password
		cn.Server = server
		cn.Database = database
		cn.Port = port
		cn.AppName = appName
		cn.CnTimeout = cnTimeout
		startTime := time.Now()
		fmt.Println(progress.Prefix("Opening database"))
		mssql.OpenCn(config.GetCnString(cn))
		fmt.Println(progress.Prefix("Reading scripts from configuration"))
		config.ReadConfig(configFile)
		fmt.Println(progress.Prefix("Executing schema scripts"))
		mssql.RunScripts(config.GetSchemaScripts())
		fmt.Println(progress.Prefix("Executing process scripts"))
		mssql.RunScripts(config.GetProcessScripts())
		elapsed := time.Since(startTime)
		fmt.Println("Total time elapsed: ", elapsed)

	},
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
	startCmd.Flags().StringVarP(&configFile, "conf", "c", "", "The configuration file")
	startCmd.Flags().StringVarP(&userName, "username", "u", "", "SQL Server user name. *Required")
	startCmd.Flags().StringVarP(&password, "password", "p", "", "SQL Server user password. *Required")
	startCmd.Flags().StringVarP(&server, "server", "s", "", "SQL Server host. *Required")
	startCmd.Flags().StringVarP(&database, "database", "d", "", "Database to work on. *Required")
	startCmd.Flags().StringVarP(&port, "port", "", "1433", "Host port number")
	startCmd.Flags().StringVarP(&cnTimeout, "timeout", "t", "14400", "Connection timeout in seconds")
	startCmd.Flags().StringVarP(&appName, "appname", "a", "go-mssql-runner", "App name to show in db calls. Useful for SQL Profiler")
}
