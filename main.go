package main

import (
	"github.com/ljahier/tunnel/pkg/client"
	"github.com/ljahier/tunnel/pkg/server"
	"github.com/spf13/cobra"
)

func runServer(cmd *cobra.Command, args []string) {
	server.Init()
}

func runClient(cmd *cobra.Command, args []string) {
	client.Init()
}

func main() {
	localServer := &cobra.Command{
		Use:   "server",
		Short: "Run server",
		Run:   runServer,
	}

	localClient := &cobra.Command{
		Use:   "client port <port>",
		Short: "Run client listening port",
		Args:  cobra.MinimumNArgs(2),
		Run:   runClient,
	}

	var rootCmd = &cobra.Command{Use: "Tunnel"}
	rootCmd.AddCommand(localServer, localClient)
	rootCmd.Execute()
}
