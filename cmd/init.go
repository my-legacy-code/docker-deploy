package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize required services",
	Long: `Initialize the required services for Docker Deploy, such as generating static assets for the web dashboard in public folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		installDependencies()
		buildDashboard()
	},
}

func installDependencies() {
	cmd := exec.Command("yarn", "install")
	cmd.Dir = "dashboard"
	cmd.Run()
}

func buildDashboard() {
	cmd := exec.Command("yarn", "build")
	cmd.Dir = "dashboard"
	cmd.Run()
}

func init() {
	rootCmd.AddCommand(initCmd)
}
