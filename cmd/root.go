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
	Long: `Runs a series of file-based SQL scripts on MS SQL Server.

A JSON configuration file is required. The configuration file should be 
in the format:

{
    "name": "A SQL Project",
    "description": "This project runs a series of scripts to manipulate data",
    "type": "Report",
    "version": "1.0.0",
    "scripts": {
        "schema": [
            "/schema/schema_script1.sql",
            "/schema/schema_script2.sql"
        ],
        "process": [
            "/process/process_script1.sql",
            "/process/process_script2.sql"
        ]
    }
}

"schema" scripts contain DDL statements such as create table or create stored
procedure statements. Files should reside under a /schema folder relative to 
the JSON configuration file.

"process" statements contain DML that run business logic. Files should reside
under a /process folder relative to the configuration file.

A "project" structure should look like:

	/project-folder
		mssqlrun.conf.json
			/schema
				schema_script1.sql
				schema_script2.sql
			/process
				process_script1.sql
				process_script2.sql

All scripts are run in the order they appear in their respective arrays. For 
example, in the "schema" array, schema_script1.sql will be run before the 
schema_script2.sql file.

WARNING: Do not place GO statements in the script files.

Connection information must be passed in via the command flags. For example:

go-mssql-runner start -c /path/configFile.json -u sqlUserName -p sqlPassword -s SQLServerHostName -d DatabaseName


`,
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
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
