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

package main

import "github.com/coolhanddev/go-mssql-runner/cmd"

/*
	TODO
		1) add "init" command to stub out config and project folder. Include example scripts.
		2) add "log" flag to start command with different levels of verbosity:
			a) verbose -- filename, duration, query result information and query text included
			b) medium -- filename, duration and query result information included
			b) default -- only filename and run duration included
		3) add "env" flag to start command to use environment variables for connection information
*/

//Version stores the app version from the git tag. Make sure to use the build command below
//
//go build -i -v -ldflags="-X main.Version=$(git describe --tags)-build$(git describe --always)"
//
//Note:  go install ignores passing of the git tag.  Run the regular build and manually cp the binary
//to the GOBIN path
var Version = "undefined"

func main() {
	cmd.AppVersion = Version
	cmd.Execute()
}
