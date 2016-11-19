# Go MSSQL Runner: a handy cli utility for running a sequence of sql script files against SQL Server 

# Overview

Using information from a JSON configuration file, this command line utility enables execution of one or a set of SQL script files.

# Platform support
Linux, OSX, Windows

# Quick Start

Navigate to the folder that has the executable.  Execute the program by issuing the start command and passing in the necessary parameters for the database connection
and the pointer to the configuration file.   

```
> go-mssql-runner start --username dbusername --password secret --server 172.0.0.1 --database yourdbname --conf /path/to/config/mssqlrun.conf.json
```

# Installation

Build the executables for the platforms supported.

Then, simply copy the appropriate executable to the destination of choice and execute.  The executables are native to the supported platforms.  There is no need
to install a runtime or a framework. There are three executables available:

1. go-mssql-runner -- executable for Linux.  
2. go-mssql-runner.exe -- executable for Windows 64 bit.
3. go-mssql-runner_osx_amd64 -- executable for OSX

For the best experience, add the location of the executable to the system path so that it can be run from any path.

# Creating a project

A project refers to the mssqlrun.conf.json file and the accompanying script files (.sql) needed to satisfy a use case. Each project should
be contained in its own folder with the following structure:

```
project-root
│   mssqlrun.conf.json│      
│
└───schema
│   │   schema_script1.sql
│   │   schema_script2.sql
│   
└───process
    │   process_script1.sql
    │   process_script1.sql
```

## The Configuration JSON file

The configuration file, mssqlrun.conf.json, controls which scripts run and the order they are run. It also contains metadata
about the project.

To create a stub, copy the contents below and save into a file named mssqlrun.conf.json.

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

**Warning**: Do not place the GO tsql keyword in the script files. If there is a need for transaction isolation, then place
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

    -a, --appname string    App name to show in db calls. Useful for SQL Profiler (default "go-mssql-runner")
    -c, --conf string       The configuration file
    -d, --database string   Database to work on. *Required
    -p, --password string   SQL Server user password. *Required
        --port string       Host port number (default "1433")
    -s, --server string     SQL Server host. *Required
    -t, --timeout string    Connection timeout in seconds (default "14400")
    -u, --username string   SQL Server user name. *Required

    ```
### Minimum required information for execution
1. username
2. password
3. server
4. database
5. complete path and filename of the configuration file.  ex. /workfolder/projectx/mssqlrun.conf.json    