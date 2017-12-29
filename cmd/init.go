// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/coolhanddev/go-mssql-runner/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a mssql.conf.json file in current directory",
	Long:  ``,
	Run:   initProj,
}

func initProj(cmd *cobra.Command, args []string) {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "01-02-2006 15:04:05", FullTimestamp: true})
	createConfig()
}

func createConfig() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	cfgFile := config.PrjConfig{
		Name:        "place project name here",
		Description: "place description here",
		Type:        "place project type here",
		Version:     "place version here",
		Scripts: config.CfgScripts{
			Schema:  []string{"schema-example-1.sql"},
			Process: []string{"process-example-1.sql"},
		},
	}
	b, _ := json.MarshalIndent(cfgFile, "", "    ")
	err = ioutil.WriteFile("mssqlrun.conf.json", b, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("mssqlrun.conf.json created.")
	fmt.Println(string(b))
}
func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
