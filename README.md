# Go MSSQL Runner 
Command line utility for running a sequence of SQL scripts.

# Overview

Using information from a JSON configuration file, this command line utility enables sequential execution of SQL script files.

# Platform support
OS: Linux, OSX, Windows 

SQL Server:  Windows and Linux
# Quick Start

Navigate to the folder that has the executable.  Execute the program by issuing the start command and passing in the necessary parameters for the database connection
and the pointer to the configuration file.   

```
$ go-mssql-runner start --username dbusername --password secret --server 172.0.0.1 --database yourdbname --conf /path/to/config/mssqlrun.conf.json
```

# Installation

Simply copy the appropriate executable to the destination of choice and execute.  The executables are native to the supported platforms.  There is no need to install a runtime or a framework. The executables are available in the release folder:

1. release/alpine-linux/go-mssql-runner
2. release/darwin/go-mssql-runner
3. release/linux/go-mssql-runner
4. release/windows/go-mssql-runner.exe

Add executable to PATH environment variable to make it available globally.

# Creating a project

A project refers to the mssqlrun.conf.json file and the accompanying script files (.sql) needed to satisfy a use case. Each project should be contained in its own folder with the following structure:

```
project-root
│   mssqlrun.conf.json      
│
└───schema
│   │   schema_script1.sql
│   │   schema_script2.sql
│   
└───process
    │   process_script1.sql
    │   process_script2.sql
```
So, creating a project involves:

1. creating a project folder 
2. creating the mssqlrun.conf.json in that folder 
3. creating the schema and process subfolders
4. creating or adding the scripts to those two subfolders 

A default mssqlrun.conf.json file can be created by issuing the command 'init' in the desired project folder

```
go-mssql-runner init
```

## The Configuration JSON file

The configuration file, mssqlrun.conf.json, controls which scripts run and the order they are run. It also contains metadata
about the project.

To create a stub, copy the contents below and save into a file named mssqlrun.conf.json or run the 'init' command as mentioned above.

```
{
    "name": "Project X",
    "description": "Aggregate and cleanse data",
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
```

The "scripts" field contains two sets of information:

### schema
This contains the list of script files that contain DDL such as create table or create stored procedure statements.
The files should reside under a /schema folder relative to the configuration file.

### process
This contains the list of script files that contain DML for the business logic.  The files should reside under a
/process folder relative to the configuration file.

**Warning**: Do not place the GO T-SQL keyword in the script files. If there is a need for transaction isolation, then place
the relevant logic in their own .sql file(s).  

#### Order of execution
Schema scripts are run first and the scripts are run in the order they appear in the list.  After the schema scripts,
the process scripts are run; they are also run in the order they appear in the list.  So, in the example schema above, 
all the scripts are executed in the following sequence:

1. schema_script1.sql
2. schema_script2.sql
3. process_script1.sql
4. process_script2.sql

# Command line
For help on which options are available, use go-mssql-runner -h or go-mssql-runner (command) -h

## Start Command
The "start" command initiates the execution.  Flags are used to specify connection and configuration needed for the program
to run the scripts against a SQL Server instance.

```    
    go-mssql-runner start [flags]

    Flags:
    -a, --appname string     App name to show in db calls. Useful for SQL Profiler (default "go-mssql-runner")
    -c, --conf string        The configuration file
    -d, --database string    Database to work on. *Required
    -e, --encrypt-cn         Encrypt SQL Server connection
    -h, --help               help for start
        --logfile string     File to write log to
        --logformat string   Format of log: JSON or text (default "text")
    -l, --loglevel string    Specifies level of verbosity for SQL log output. See start command help (above) for details (default "0")
    -p, --password string    SQL Server user password. *Required
        --port string        Host port number (default "1433")
    -s, --server string      SQL Server host. *Required
    -t, --timeout string     Connection timeout in seconds (default "14400")
    -u, --username string    SQL Server user name. *Required
```

### Minimum information for execution
1. username -- in Windows, if username is not given, integrated authentication will be attempted
2. password
3. server
4. database
5. complete path and filename of the configuration file.  ex. /workfolder/projectx/mssqlrun.conf.json

## Environment Variables

The following environment variables can be set to store the minimum information.  If no flags are set on the command, then the program will look at these for values.  Flags override environment values.

**GOSQLR_CONFIGFILE**

**GOSQLR_USERNAME**

**GOSQLR_PASSWORD**

**GOSQLR_SERVER**

**GOSQLR_DATABASE**

# Example usage

Below is an example for Windows, executing and passing command line options.  The executable either exists in the \MyDBProject folder or it is
globally made available via the PATH environment variable. 
```
c:\MyDBProject>go-mssql-runner -u sa -p secret -s localhost -d adventureworks2012 -c ./mssqlrun.conf.json 
```
Setting the environment variables first in the shell and then executing the utility
```
c:\MyDBProject>set GOSQLR_CONFIGFILE=./mssqlrun.conf.json
c:\MyDBProject>set GOSQLR_USERNAME=sa
c:\MyDBProject>set GOSQLR_PASSWORD=secret
c:\MyDBProject>set GOSQLR_SERVER=localhost
c:\MyDBProject>set GOSQLR_DATABASE=adventureworks2012
c:\MyDBProject>go-mssql-runner start
```
During command line execution, the program can be run by invoking start as shown above.  One or more of the values can be overriden as shown in 
the example below where the config location is overriden by the value passed in the -c flag:

```
c:\MyDBProject>go-mssql-runner start -c /SomeOtherProjectFolder/mssqlrun.conf.json
``` 

Linux and OSX usage are similar. Assuming the executable is made globally available by adding its location to the PATH via an init script
```
$ go-mssql-runner -u sa -p secret -s localhost -d adventureworks2012 -c ~/MyDBProject/mssqlrun.conf.json
```
Setting the environment variables to last only the lifetime of the terminal session.
```
$ export GOSQLR_CONFIGFILE=~/MyDBProject/mssqlrun.conf.json
$ export GOSQLR_USERNAME=sa
$ export GOSQLR_PASSWORD=secret
$ export GOSQLR_SERVER=localhost
$ export GOSQLR_DATABASE=adventureworks2012
$ go-mssql-runner start
```

```
$ go-mssql-runner start -c ~/SomeOtherProjectFolder/mssqlrun.conf.json
```
# Docker container
The included Dockerfile serves as an example on how a project image can be built.   Replace the example folders with your preferred project folder. For a light container, use the Alpine image and the alpine build of the executable.

To run a persistent container:
```
$ docker container run --name containername go-mssql-runner-imagename start -u someuser -p somepass -s someserverip -d somedbname -c someconfigfile.json --logformat JSON
```
To access the container log and save to file
```
docker container logs containername > log.txt 
```
To run an ephemereal container and save log to file
```
docker container run -i --rm go-mssql-runner-imagename start -u someuser -p somepass -s someserverip -d somedbname -c someconfigfile.json --logformat JSON  2>&1 | tee -a log.txt
```

# Tips

* Log all the time.
* Use the command history of your favorite shell to rerun the command. Typically accessed by up or down arrow key.
* Get detailed account of what ran using the different log level in the -l flag of start command.
* Save the screen output to a text file by specifiing the --logfile flag of the start command.
* Use the JSON option on the --logformat flag to outout the log in JSON format and make it easy to parse.
* Although scripts are run sequentially, concurrent operation can be accomplished via scripting or cli techniques. Separate operations that can be run concurrently into their own projectsFor example in Bash, commands can be run in parallel using '&': 
* Build Docker images to encapsulate versions of your project.  For example, create an image for dev/test and another image for production.
* Docker containers can be run repeatedly. Command parameters need not be specified each time. Just run "docker start <containerName>" to rerun projects.


```
$ command1 & command2 
``` 
   

# Roadmap
* script encryption
* embedded projects
* Slack integration
* support for cloud logging (AWS CloudWatch)