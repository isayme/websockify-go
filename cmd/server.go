package cmd

import (
	"github.com/isayme/websockify-go/cmd/server"
	"github.com/spf13/cobra"
)

var serverOptions = server.ServerOptions{}

func init() {
	serverOptions.Listen = serverCmd.Flags().String("listen", ":6080", "Port for proxy/webserver to listen on")
	serverOptions.Vnc = serverCmd.Flags().String("vnc", "localhost:5900", "VNC server host:port proxy target")
	serverOptions.Web = serverCmd.Flags().String("web", "./", "Path to web files (e.g. vnc.html)")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server of websockify",
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(serverOptions)
	},
}
