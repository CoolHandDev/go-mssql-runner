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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-mssql-runner",
	Short: "A script runner for MS SQL Server",
	Long: `This utility runs a series of SQL scripts against MS SQL Server.

The utility relies on a JSON configuration file, named mssqlrun.conf.json,
to find out which scripts to run.  The configuration file should be in
the format specified below:

{
    "name": "A SQL Project",
    "description": "This project runs a series of scripts to manipulate data",
    "type": "Report",
    "version": "1.0.0",
    "scripts": {
        "schema": [
            "schema_script1.sql",
            "schema_script2.sql"
        ],
        "process": [
            "process_script1.sql",
            "process_script2.sql"
        ]
    }
}

"schema" scripts contain DDL statements such as create table or create stored
procedure statements. "process" statements contain DML that run business logic.
All scripts are run in the order they appear in their respective arrays. For 
example, in the "schema" array, schema_script1.sql will be run before the 
schema_script2.sql file.


In addition to specifiying a configuration, the proper connection information
must be passed in via the command flags. For example:

go-mssql-runner start -c /PathToConfig/ -u sqlUserName -p sqlPassword -s SQLServerHostName -d DatabaseName


`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-mssql-runner.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".go-mssql-runner") // name of config file (without extension)
	viper.AddConfigPath("$HOME")            // adding home directory as first search path
	viper.AutomaticEnv()                    // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
