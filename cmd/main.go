package main

import (
	"github.com/c-bata/goptuna/cmd/feic"
	"github.com/spf13/cobra"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	version  string
	revision string

	rootCmd = &cobra.Command{
		Use:   "goptuna",
		Short: "A command line interface for Goptuna",
	}
)

func main() {
	//rootCmd.AddCommand(createstudy.GetCommand())
	//rootCmd.AddCommand(deletestudy.GetCommand())
	//rootCmd.AddCommand(dashboard.GetCommand())
	//if version != "" && revision != "" {
	//	rootCmd.Version = fmt.Sprintf("%s (rev: %s)", version, revision)
	//}
	//err := rootCmd.Execute()
	//if err != nil {
	//	rootCmd.PrintErrln(err)
	//	os.Exit(1)
	//}
	cmd := feic.GetCommand()
	err := cmd.Execute()
	if err != nil {
		cmd.PrintErrln(err)
		os.Exit(1)
	}
}
