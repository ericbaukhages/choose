package cmd

import (
	"ebaukhages/choose/choose"
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new tmux session",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config := choose.Config{
			Location: "/Users/ebaukhages/Documents/scripts/tmux.sessions.log",
		}
		config.Parse()

		var (
			name string
			path string
		)

		if len(args) == 2 {
			name = args[0]
			path = args[1]
		} else {
			fmt.Println("Insufficient arguments.")
			return
		}

		err := config.Add(name, path)
		if err != nil {
			fmt.Printf("New session could not be added: %v\n", err)
		}

		err = config.Save()
		if err != nil {
			fmt.Printf("Project could not be saved: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
