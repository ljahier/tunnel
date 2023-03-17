package main

import (
	"strconv"

	"github.com/ljahier/tunnel/internal/client"
	"github.com/ljahier/tunnel/internal/server"
	"github.com/spf13/cobra"
)

func runServer(cmd *cobra.Command, args []string) {
	port, err := strconv.Atoi(cmd.Flag("port").Value.String())

	if err != nil {
		panic("You must enter a number as port")
	}
	server.Init(port)
}

func runClient(cmd *cobra.Command, args []string) {
	port, err := strconv.Atoi(cmd.Flag("port").Value.String())
	if err != nil {
		panic("You must enter a number as port")
	}
	server := cmd.Flag("server").Value.String()
	value := cmd.Flag("value").Value.String()

	client.Init(server, port, value)
}

/*
	How tunnel works ?
	- the client make an http request to ask subdomain and to open new connection
	- the server expose subdomain and redirect all incomming request from this subdomain to client inside already open tcp conn
*/

func main() {
	// TODO(lucas): split cli command into another src file
	localServer := &cobra.Command{
		Use:   "server",
		Short: "Server cli",
	}

	localServerStart := &cobra.Command{
		Use:   "start",
		Short: "Run client listening port",
		Run:   runServer,
	}

	// localServerPort := &cobra.Command{
	// 	Use:   "port",
	// 	Short: "Run client on specified port",
	// 	Args:  cobra.MaximumNArgs(1),
	// 	Run:   runServer,
	// }

	localClient := &cobra.Command{
		Use:   "client",
		Short: "Run client",
		Run:   runClient,
	}

	var rootCmd = &cobra.Command{Use: ""}

	rootCmd.AddCommand(localServer, localClient)
	localServer.AddCommand(localServerStart)
	localServerStart.PersistentFlags().Int("port", 3000, "Server port")
	localClient.PersistentFlags().Int("port", 8081, "Port which are listenning for your local web service")
	localClient.PersistentFlags().String("server", "127.0.0.1:3000", "Tunneling server url")
	localClient.PersistentFlags().String("value", "1", "Value")
	rootCmd.Execute()
}
