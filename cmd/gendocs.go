package cmd
import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var genDocsCmd = &cobra.Command{
	Use:   "gen-docs",
	Short: "Generate documentations for Docker Deploy CLI",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) {
		doc.GenMarkdownTree(rootCmd, "docs/markdown")
	},
}

func init() {
	rootCmd.AddCommand(genDocsCmd)
}
