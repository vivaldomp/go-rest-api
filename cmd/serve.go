package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vivaldomp/go-rest-api/routers"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Long:  `Starts application and serves the configured system`,
	Run: func(cmd *cobra.Command, args []string) {
		server, err := routers.NewServer()
		cobra.CheckErr(err)
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
