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

	"github.com/coolhanddev/go-mssql-runner/pkg/config"
	"github.com/spf13/cobra"
)

var configFile string
var userName string
var password string
var server string
var database string
var port int
var cnTimeout int
var encrypt bool
var appName string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start running scripts",
	Long: `

The start command kicks off the execution of the scripts listed in 
the configuration json file specified in the --config flag. The 
proper connection information to MS SQL Server must be passed in.

`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("start called", configFile, args)
		fmt.Println(config.GetCnString("sa", "dev", "localhost", "adventureworks2012"))
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

	startCmd.Flags().StringVarP(&configFile, "conf", "c", "", "Location of project.conf.json configuration file")
	startCmd.Flags().StringVarP(&userName, "username", "u", "", "SQL Server user name")
	startCmd.Flags().StringVarP(&password, "password", "p", "", "SQL Server user password")
	startCmd.Flags().StringVarP(&server, "server", "s", "", "SQL Server host")
	startCmd.Flags().StringVarP(&database, "database", "d", "", "Database to work on")
	startCmd.Flags().IntVarP(&port, "port", "", 1433, "Host port number")
	startCmd.Flags().IntVarP(&cnTimeout, "timeout", "t", 14400, "Connection timeout in seconds")
	startCmd.Flags().BoolVarP(&encrypt, "encrypt", "e", false, "Encrypt the connection true/false (default false)")
	startCmd.Flags().StringVarP(&appName, "appname", "a", "go-mssql-runner", "App name to show in db calls. Useful for SQL Profiler")

}
