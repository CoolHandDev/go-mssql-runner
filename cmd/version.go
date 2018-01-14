package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//AppVersion stores version information from main
var AppVersion string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version",
	Long:  `Show the version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-sql-runner:", AppVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
