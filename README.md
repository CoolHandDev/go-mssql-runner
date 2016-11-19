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

Simply copy the appropriate executable to the destination of choice and execute.  The executables are native to the supported platforms.  There is no need
to install a runtime or a framework. There are three executables available

1. go-mssql-runner -- executable for Linux.  
2. go-mssql-runner.exe -- executable for Windows 64 bit.
3. go-mssql-runner_osx_amd64 -- executable for OSX

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
│   │
│   
└───process
    │   process_script1.sql
    │   process_script1.sql
```