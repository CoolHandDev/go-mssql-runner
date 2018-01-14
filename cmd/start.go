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
var encryptCn bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start running scripts",
	Long:  startCmdLongDesc(),
	Run:   start,
}

func init() {
	RootCmd.AddCommand(startCmd)

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
	startCmd.Flags().BoolVarP(&encryptCn, "encrypt-cn", "e", false, "Encrypt SQL Server connection")
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
	cn.Encrypt = encryptCn
	startTime := time.Now()
	initLogging()
	defer logFileName.Close()
	newDbPool(cn)
	loadConfig()
	execScripts()
	elapsed := time.Since(startTime)
	if logToFile != "" {
		log.WithFields(log.Fields{"log_file": logToFile}).Info("Log file created")
	}
	log.Info("Total time elapsed", "=", elapsed)
}

func newDbPool(c config.MssqlCn) {
	log.WithFields(log.Fields{"server": cn.Server, "database": cn.Database}).Info("opening database")
	mssql.NewPool(config.GetCnString(c))
}

func initLogging() {
	//we want to log both to stdout and to a file
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
}

func loadConfig() {
	log.WithFields(log.Fields{"config_file": configFile}).Info("loading configuration")
	config.ReadConfig(configFile)
}

func execScripts() {
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
}

func startCmdLongDesc() string {
	return `

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
	
	`
}
