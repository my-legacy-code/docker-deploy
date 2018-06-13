package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/Teamyapp/docker-deploy/core"
)

var configFilename string
var port string

var rootCmd = &cobra.Command{
	Use: "docker-deploy",
	Short: "Docker Deploy is a easy to use container manager",
	Long: `A Fast and Flexible Container Manager built with love by byliuyang and friends in Go.
Complete documentation is available at https://github.com/Teamyapp/docker-deploy`,
	Run: func(cmd *cobra.Command, args []string) {
		core.LaunchServer(configFilename, port)
	},
}

func init()  {
	rootCmd.PersistentFlags().StringVarP(&configFilename, "config", "c", "testdata/config_test.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "3000", "port the server listen on")
}

func Execute()  {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}