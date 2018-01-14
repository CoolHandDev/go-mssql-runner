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
	b, err := json.MarshalIndent(cfgFile, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("mssqlrun.conf.json", b, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("mssqlrun.conf.json created.")
	fmt.Println(string(b))
}
func init() {
	RootCmd.AddCommand(initCmd)
}
