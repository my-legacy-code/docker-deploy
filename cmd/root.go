package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	//"github.com/Teamyapp/docker-deploy/core"
)

var rootCmd = &cobra.Command{
	Use: "docker-deploy",
	Short: "Docker Deploy is a easy to use container manager",
	Long: `A Fast and Flexible Container Manager build with love by byliuyang and friends in Go.
Complete documentation is available at https://github.com/Teamyapp/docker-deploy`,
	Run: func(cmd *cobra.Command, args []string) {
		//cmd.flag
		//core.LaunchServer(args)
		fmt.Println("Good")
	},
}

func init()  {
	fmt.Println("Good")
}

func Execute()  {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}