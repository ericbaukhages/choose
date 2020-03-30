package cmd

import (
	"fmt"

	"github.com/ericbaukhages/choose/choose"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all stored sessions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configFileName, err := homedir.Expand("~/.tmux.sessions.log")
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}

		config := choose.Config{
			Location: configFileName,
		}
		config.Parse()

		err = config.Print()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
