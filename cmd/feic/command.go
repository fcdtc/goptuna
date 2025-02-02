package feic

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"

	"github.com/c-bata/goptuna/dashboard"
	"github.com/c-bata/goptuna/rdb.v2"
	"github.com/spf13/cobra"
)

// GetCommand returns the cobra's command for create-study sub-command.
func GetCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "dashboard",
		Short: "Launch web dashboard",
		Run: func(cmd *cobra.Command, args []string) {
			database, err := cmd.Flags().GetString("record")
			db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			storage := rdb.NewStorage(db)
			server, err := dashboard.NewServer(storage)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			hostName, err := cmd.Flags().GetString("ip")
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
			port, err := cmd.Flags().GetString("port")
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
			addr := fmt.Sprintf("%s:%s", hostName, port)

			cmd.Printf("Started to serve at http://%s\n", addr)
			if err := http.ListenAndServe(addr, server); err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
		},
	}
	command.Flags().StringP("record", "r", "", "database name")
	command.Flags().StringP("ip", "i", "127.0.0.1", "IP address")
	command.Flags().StringP("port", "p", "8000", "port number")
	return command
}
